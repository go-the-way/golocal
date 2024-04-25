package golocal

import "sync"

type defaultContext[T any] struct {
	m map[string]T
	*sync.RWMutex
}

func newDefaultContext[T any]() Context[T] {
	return &defaultContext[T]{map[string]T{}, &sync.RWMutex{}}
}

func (r *defaultContext[T]) localSet(data T)    { r.Lock(); defer r.Unlock(); r.m[GoId()] = data }
func (r *defaultContext[T]) localGet() (data T) { r.RLock(); defer r.RUnlock(); return r.m[GoId()] }
func (r *defaultContext[T]) Set(data T)         { r.localSet(data) }
func (r *defaultContext[T]) Get() (data T)      { return r.localGet() }
