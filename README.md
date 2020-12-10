# wui [![GoDoc](https://godoc.org/github.com/golangee/wui?status.svg)](http://godoc.org/github.com/golangee/wui)
Package wui provides a Renderable abstraction for manipulating the html dom. It provides three base building
blocks: Node, Modifier and Component. A Node allocates dom elements whereas a Modifier changes it by either
modifying attributes are using other Node instances to append elements. Lastly, a Component returns a Node
to allocate an Element but also connects it with a lifecycle, as long as it is *in use*. Detaching
an Element from its parent will release it (mostly listeners connected to a Component). Subsequent Component.Render 
calls will allocate another Element and attach until released (and/or re-rendered). There is no shadow or virtual
dom to improve efficiency. To compensate and even surpass the performance of state-of-the-art vdom driven 
frameworks, one can attach modifiers with a live element, without replacing any elements or performing
expensive delta-dom computations.

## What is a *wui*?
wui is an abbreviation for *web user interface*. It is the successor of *forms* which itself was not abstract enough
to justify that name. *wui* is a very small core with the most important functions to drive a *wasm* web app.

See also https://github.com/golangee/html for a convenience library to build your own Component.

## Roadmap
Stabilize the API. Subsume it under a future implementation and technology independent *forms* project.
