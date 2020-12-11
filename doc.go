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

// Package gotrino provides a Renderable abstraction for manipulating the html dom. It provides three base building
// blocks: Node, Modifier and Component. A Node allocates dom elements, a Modifier changes it by either
// modifying attributes are using other Node instances to append elements or a Component which returns a Node
// to allocate and modify an Element but also connects it with a lifecycle, as long as it is alive. Detaching
// an Element will release it (mostly listeners connected to a Component). Subsequent Component.Render calls
// will allocate another Element and attach until released (and/or re-rendered).
package gotrino
