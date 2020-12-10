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

package wui

import "github.com/golangee/dom"

// A Node gives access to a dom.Element. Usually the (render) Node allocates a new
// Element, which can be attached.
type Node interface {
	// Element returns the wrapped Element, usually a new instance is allocated.
	Element() dom.Element
	Renderable
}

// NodeFunc is a func type which allows usage as a Node. It is the only way to create such
// an implementation.
type NodeFunc func() dom.Element

// Element returns the wrapped Element, usually a new instance is allocated.
func (f NodeFunc) Element() dom.Element {
	return f()
}

// nodeOrModifierOrComponent is our private marker contract.
func (f NodeFunc) nodeOrModifierOrComponent() {
}
