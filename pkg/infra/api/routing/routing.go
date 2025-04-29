package routing

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Middleware is an alias for fiber.Handler
type Middleware = fiber.Handler

// Route struct that wraps *fiber.App
type Router struct {
	app        *fiber.App
	groupPath  string
	middleware []Middleware
}

// NewRouter initializes a new Fiber router
func NewRouter() *Router {
	return &Router{
		app: fiber.New(),
	}
}

// joinPath ensures proper route formatting
func (r *Router) joinPath(path string) string {
	if len(path) == 0 || path[0] != '/' {
		panic("Path should start with '/' in '" + path + "'.")
	}
	return r.groupPath + path
}

// Group creates a sub-router with shared middleware
func (r *Router) Group(path string, m ...Middleware) *Router {
	return &Router{
		app:        r.app,
		groupPath:  r.joinPath(path),
		middleware: append(r.middleware, m...),
	}
}

// Use applies global middleware
func (r *Router) Use(m ...Middleware) {
	r.middleware = append(r.middleware, m...)
}

// Handler registers a route with middleware
func (r *Router) Handle(method, path string, handler fiber.Handler, middleware ...Middleware) {
	fullPath := r.joinPath(path)

	// Merge global middleware with route-specific middleware
	handlers := append(r.middleware, middleware...)
	handlers = append(handlers, handler)

	// Register route with all middleware
	r.app.Add(method, fullPath, handlers...)
}

// Shorcut methods for HTTP verbs
func (r *Router) GET(path string, handler fiber.Handler, middleware ...Middleware) {
	r.Handle(fiber.MethodGet, path, handler, middleware...)
}

func (r *Router) POST(path string, handler fiber.Handler, middleware ...Middleware) {
	r.Handle(fiber.MethodPost, path, handler, middleware...)
}

func (r *Router) PUT(path string, handler fiber.Handler, middleware ...Middleware) {
	r.Handle(fiber.MethodPut, path, handler, middleware...)
}

func (r *Router) DELETE(path string, handler fiber.Handler, middleware ...Middleware) {
	r.Handle(fiber.MethodDelete, path, handler, middleware...)
}

// ServeStatic serves static files from a directory
func (r *Router) ServeStatic(path, root string) {
	r.app.Static(path, root)
}

// Start the Fiber server
func (r *Router) Start(address string) error {
	return r.app.Listen(address)
}

// Gracefully shut down the Fiber app
func (r *Router) Shutdown(ctx context.Context) error {
	return r.app.ShutdownWithContext(ctx)
}

func (r *Router) ListenToAddress(address string) error {
	return r.app.Listen(address) // Start the Fiber app
}
