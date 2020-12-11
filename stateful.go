// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gotrino

import "sync"

// Handle represents manual lifecycle manager to a registered Stateful observer callback.
type Handle struct {
	parent *Stateful
	idx    int
}

// Release detached the handle from its parent, so that the registered function can
// be garbage collected, as long as there are no other references. A call is idempotent.
func (h Handle) Release() {
	if h.parent == nil {
		return
	}

	h.parent.lock.Lock()
	defer h.parent.lock.Unlock()

	h.parent.observers[h.idx] = nil
}

// Stateful is an observable helper which can notify all registered callbacks
// when calling Invalidate. The usage is concurrency and recursive safe.
type Stateful struct {
	observers []func()   // we only append, invalid observers will be set to nil
	lock      sync.Mutex // we ever need this very short, no deadlocks possible
}

// Invalidate will invoke all observers which have not been released yet.
// Observers are free to release their Handle or to register new observers.
func (c *Stateful) Invalidate() {
	c.lock.Lock()
	length := len(c.observers)
	c.lock.Unlock()

	for i := 0; i < length; i++ {
		c.lock.Lock()
		observer := c.observers[i]
		c.lock.Unlock()

		if observer != nil {
			observer()
		}
	}
}

// Observe registers the given function to be called the next time Invalidate is called.
func (c *Stateful) Observe(f func()) Handle {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.observers = append(c.observers, f)

	return Handle{parent: c, idx: len(c.observers) - 1}
}
