package golocal

import "sync"

type DefaultContext[T any] struct {
	m map[string]T
	*sync.RWMutex
}

func NewDefaultContext[T any]() Context[T] {
	return &DefaultContext[T]{map[string]T{}, &sync.RWMutex{}}
}

func (d *DefaultContext[T]) localSet(data T)    { d.Lock(); defer d.Unlock(); d.m[GoId()] = data }
func (d *DefaultContext[T]) localGet() (data T) { d.RLock(); defer d.RUnlock(); return d.m[GoId()] }
func (d *DefaultContext[T]) Set(data T)         { d.localSet(data) }
func (d *DefaultContext[T]) Get() (data T)      { return d.localGet() }
