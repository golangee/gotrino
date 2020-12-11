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

import "github.com/golangee/dom"

// Renderable is a marker interface which identifies one of our three kinds of
// DOM manipulation primitives. This must be one of Component, Node or Modifier.
type Renderable interface {
	// nodeOrModifierOrComponent is our private marker contract.
	nodeOrModifierOrComponent()
}

// A Component interface is currently only implementable by embedding a View. You may ask "if there is just one
// implementation, why would you need an interface?". The answer is, because we have as many implementations,
// as you create, however, only a part for the Component (so a View is not yet a Component) contract can
// be introduced by embedding a View. We require to rely on dynamic polymorphic method dispatching, which
// can only be achieved by using interfaces.
type Component interface {
	// Render returns a view root Node.
	Render() Node
	// Observe registers with the component which notifies for changes.
	Observe(f func()) Handle
	Renderable

	getPostModifiers() []Modifier
	setPostModifiers(mods ...Modifier)
}

// RenderBody clears the body of the page and applies the given Renderable.
func RenderBody(c Renderable) {
	body := dom.GetWindow().Document().Body()
	body.Clear()
	WithElement(body, c).Element()
}
