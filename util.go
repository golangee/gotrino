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

import (
	"fmt"
	"reflect"

	"github.com/golangee/dom"
	"github.com/golangee/property"
)

// Join is a convenience operator to merge 1+n renderables into a common slice.
func Join(r Renderable, other ...Renderable) []Renderable {
	tmp := make([]Renderable, 0, len(other)+1)

	tmp = append(tmp, r)
	tmp = append(tmp, other...)

	return tmp
}

// Element creates a new dom Element using a delayed Node allocation. See also ElementNS to create an Element
// in a specific namespace.
func Element(name string, rm ...Renderable) Node {
	return WithElement(dom.GetWindow().Document().CreateElement(name), rm...)
}

// ElementNS create a new Element with a specific namespace.
// Normal HTML elements do not need a namespace, however some requires to work properly, e.g. like SVG declarations.
func ElementNS(namespace, name string, rm ...Renderable) Node {
	return WithElement(dom.GetWindow().Document().CreateElementNS(namespace, name), rm...)
}

// WithElement applies the given renderables into the context of the element. This usually either means
// appending the Node.Element from a Node or Component as a child to the given element or by modifying
// the given Element (which in turn may recursively append dom.Element's).
func WithElement(elem dom.Element, rm ...Renderable) Node {
	return NodeFunc(func() dom.Element {
		for _, e := range rm {
			switch t := e.(type) {
			case Node:
				elem.AppendElement(t.Element())
			case Modifier:
				t.Modify(elem)
			case Component:
				x := t.Render().Element()
				for _, modifier := range t.getPostModifiers() {
					modifier.Modify(x)
				}
				var observer func()
				var xHandle Handle
				observer = func() {
					x.Release()
					xHandle.Release()
					newElem := t.Render().Element()
					for _, modifier := range t.getPostModifiers() {
						modifier.Modify(newElem)
					}
					x = x.ReplaceWith(newElem)
					xHandle = t.Observe(observer)
					x.AddReleaseListener(func() {
						xHandle.Release()
					})
				}
				xHandle = t.Observe(observer)

				x.AddReleaseListener(func() {
					xHandle.Release()
				})
				elem.AppendElement(x)
			case nil:
				// this makes optional sub-components easier
			default:
				panic(fmt.Sprintf("the type '%s' must be either a Node, a Modifier or a Component. "+
					"Did you forget to add the Render method?", reflect.TypeOf(t).String()))
			}
		}

		return elem
	})
}

// If only evaluates the flag once and can not be changed afterwards. It is useful in non-components
// or if properties are not needed, because a full rendering will be done anyway. See also IfCond for
// a more efficient way of changing properties.
func If(flag bool, pos, neg Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if flag {
			if pos != nil {
				WithElement(e, pos).Element()
			}
		} else {
			if neg != nil {
				WithElement(e, neg).Element()
			}
		}
	})
}

// IfCond applies the given positive and negative modifiers in-place, without causing
// an entire re-rendering, if the property changes. This improves performance
// a lot. See also If.
func IfCond(p *property.Bool, pos, neg Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if p.Get() {
			if pos != nil {
				pos.Modify(e)
			}
		} else {
			if neg != nil {
				neg.Modify(e)
			}
		}

		h := p.Observe(func(old, new bool) {
			if new {
				pos.Modify(e)
			} else {
				neg.Modify(e)
			}
		})

		e.AddReleaseListener(h.Release)
	})
}

// With post-modifies the given Renderable for each future rendering.
func With(r Renderable, mods ...Modifier) Renderable {
	switch t := r.(type) {
	case Component:
		return WithComponent(t, mods...)
	case Node:
		return NodeFunc(func() dom.Element {
			elem := t.Element()
			tmp := make([]Renderable, 0, len(mods))
			for _, mod := range mods {
				tmp = append(tmp, mod)
			}
			WithElement(elem, tmp...)

			return elem
		})
	case Modifier:
		return ModifierFunc(func(e dom.Element) {
			for _, mod := range mods {
				mod.Modify(e)
			}
		})
	default:
		panic(reflect.TypeOf(r).String())
	}
}

// InsideDom invokes the callback for each invocations. Be careful not to leak the element. Think twice before
// using and consider dom.Element#AddReleaseListener.
func InsideDom(f func(e dom.Element)) Modifier {
	return ModifierFunc(func(e dom.Element) {
		f(e)
	})
}

// ForEach is very useful to simply represent a dynamic loop in a functional way, especially for appending
// elements.
func ForEach(len int, f func(i int) Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for i := 0; i < len; i++ {
			x := f(i)
			WithElement(e, x).Element()
		}
	})
}

// Yield is a convenience operator to apply or insert multiple renderables as one, especially useful if one
// ever needs to evaluate or append multiple renderables without a container.
func Yield(r ...Renderable) Renderable {
	return ModifierFunc(func(e dom.Element) {
		for _, renderable := range r {
			WithElement(e, renderable).Element()
		}
	})
}

// WithComponent post-modifies the given Component for each future rendering.
func WithComponent(r Component, mods ...Modifier) Component {
	r.setPostModifiers(mods)

	return r
}

// Modifiers aggregates multiple Modifier into a single modifier delegate.
func Modifiers(m ...Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for _, modifier := range m {
			modifier.Modify(e)
		}
	})
}
