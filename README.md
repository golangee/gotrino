# gotrino [![Go Reference](https://pkg.go.dev/badge/github.com/golangee/gotrino.svg)](https://pkg.go.dev/github.com/golangee/gotrino)
Package gotrino provides a Renderable abstraction for manipulating the html dom. It provides three base building
blocks: Node, Modifier and Component. A Node allocates dom elements whereas a Modifier changes it by either
modifying attributes are using other Node instances to append elements. Lastly, a Component returns a Node
to allocate an Element but also connects it with a lifecycle, as long as it is *in use*. Detaching
an Element from its parent will release it (mostly listeners connected to a Component). Subsequent Component.Render 
calls will allocate another Element and attach until released (and/or re-rendered). There is no shadow or virtual
dom to improve efficiency. To compensate and even surpass the performance of state-of-the-art vdom driven 
frameworks, one can attach modifiers with a live element, without replacing any elements or performing
expensive delta-dom computations.

## What is a *gotrino*?
A *gotrino* is a fictional elementary particle. The successor was named *forms* which itself was not abstract 
enough to justify that name. *gotrino* itself is a very small core with the most battery functions to 
drive a *wasm* web app.

See also https://github.com/golangee/gotrino-html for a convenience library to build your own Components.

## Roadmap
Stabilize the API. Subsume it under a future implementation and technology independent *forms* project.
