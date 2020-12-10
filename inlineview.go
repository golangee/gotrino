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

// InlineView is a helper to define views in an abstract way, without the
// need to declare a (private) struct with the according method.
type InlineView struct {
	View
	f func() Node
}

// NewInlineView creates an InlineView which invokes the given closure on Render.
func NewInlineView(render func() Node) *InlineView {
	return &InlineView{f: render} //nolint:exhaustivestruct
}

// Render returns a view root Node.
func (v InlineView) Render() Node {
	return v.f()
}
