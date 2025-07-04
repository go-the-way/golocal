// Copyright 2025 golocal Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
