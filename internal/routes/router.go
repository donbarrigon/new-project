package routes

import "net/http"

// Create a new router instance
var router = http.NewServeMux()

type route struct {
	Path    string
	Handler http.HandlerFunc
	Name    string
}

func Setup() http.Handler {

	// Add public routes to router
	for _, r := range healthCheck {
		router.HandleFunc(r.Path, r.Handler)
	}
	return router
}

// registra las rutas en el router
func registerRoutes(prefix string, middleware []func(http.HandlerFunc) http.HandlerFunc, handler []route) {
	for _, r := range handler {
		router.HandleFunc(prefix+r.Path, use(middleware, r.Handler))
	}
}

// Add middleware here, if needed.
func use(middleware []func(http.HandlerFunc) http.HandlerFunc, handler http.HandlerFunc) http.HandlerFunc {
	// Aplica los middlewares de forma anidada
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}

// Create route GET
func handleGet(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route POST
func handlePost(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route PUT
func handlePut(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route DELETE
func handleDelete(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route PATCH
func handlePatch(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route HEAD
func handleHead(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route OPTIONS
func handleOptions(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route TRACE
func handleTrace(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodTrace {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// Create route CONNECT
func handleConnect(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodConnect {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
