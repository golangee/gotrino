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

// Modifier is Renderable which changes attributes or contents of the given dom.Element.
type Modifier interface {
	// Modify applies its changes to the given element.
	Modify(e dom.Element)
	Renderable
}

// ModifierFunc is a func type which allows usage as a Modifier. It is the only way to create such
// an implementation.
type ModifierFunc func(e dom.Element)

// Modify applies its changes to the given element.
func (f ModifierFunc) Modify(e dom.Element) {
	f(e)
}

// nodeOrModifierOrComponent is our private marker contract.
func (f ModifierFunc) nodeOrModifierOrComponent() {
}
