package aero

import (
	"net/http"
	"os"
)

// Router is a high-performance router.
type Router struct {
	get     tree
	post    tree
	delete  tree
	put     tree
	patch   tree
	head    tree
	connect tree
	trace   tree
	options tree
}

// Add registers a new handler for the given method and path.
func (router *Router) Add(method string, path string, handle Handle) {
	tree := router.selectTree(method)
	tree.add(path, handle)
}

// Find returns the handle for the given route.
// This is only useful for testing purposes.
// Use Lookup instead.
func (router *Router) Find(method string, path string) Handle {
	c := Context{}
	router.Lookup(method, path, &c)
	return c.handler
}

// Lookup finds the handle and parameters for the given route
// and assigns them to the given context.
func (router *Router) Lookup(method string, path string, ctx *Context) {
	tree := router.selectTree(method)

	// Fast path for the root node
	if tree.prefix == path {
		ctx.handler = tree.data
		return
	}

	tree.find(path, ctx)
}

// Print shows a pretty print of the routes.
func (router *Router) Print(method string) {
	tree := router.selectTree(method)
	tree.PrettyPrint(os.Stdout)
}

// selectTree returns the tree by the given HTTP method.
func (router *Router) selectTree(method string) *tree {
	switch method {
	case http.MethodGet:
		return &router.get
	case http.MethodPost:
		return &router.post
	case http.MethodDelete:
		return &router.delete
	case http.MethodPut:
		return &router.put
	case http.MethodPatch:
		return &router.patch
	case http.MethodHead:
		return &router.head
	case http.MethodConnect:
		return &router.connect
	case http.MethodTrace:
		return &router.trace
	case http.MethodOptions:
		return &router.options
	default:
		return nil
	}
}
