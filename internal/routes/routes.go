package routes

import (
	"log"
	"net/http"

	"github.com/erespereza/new-project/internal/controller"
	"github.com/erespereza/new-project/internal/middleware"
)

// constantes para poder escribir rutas bonitas
const (
	get     = http.MethodGet     // Solicita la representación de un recurso.
	post    = http.MethodPost    // Envía datos al servidor para procesarlos.
	put     = http.MethodPut     // Reemplaza completamente el recurso con los datos proporcionados.
	delete  = http.MethodDelete  // Elimina un recurso.
	patch   = http.MethodPatch   // Aplica modificaciones parciales a un recurso.
	head    = http.MethodHead    // Similar a GET, pero solo recupera los encabezados.
	options = http.MethodOptions // Describe las opciones de comunicación para el recurso.
	connect = http.MethodConnect // Establece un túnel para comunicaciones (usado principalmente con proxies HTTPS).
	trace   = http.MethodTrace   // Realiza una prueba de bucle de retorno para el recurso.
)

// Create a new router instance
var router = http.NewServeMux()

// Registra las rutas en el router
type Route struct {
	Path    string
	Methods []string
	Handler controller.ControllerFunc
	Name    string
}

func NewRouter() *http.ServeMux {

	//rutas web
	HandleFunc("", webPublic)
	HandleFunc("", webPrivate, middleware.Logger, middleware.Request)

	//rutas api
	HandleFunc("/api/v1", apiPublic)
	HandleFunc("/api/v1", apiPrivate, middleware.Logger, middleware.Request)

	//este print esta ahi nada mas para que no salten un monton de errores en las constantes
	log.Println("Rutas establecidas, metodos habilidatos: " + get + ", " + post + ", " + put + ", " + delete + ", " + patch + ", " + head + ", " + options + ", " + connect + ", " + trace)

	return router
}

// AllowMethods funcion auxiliar para crear un slice de strings y que se vea bonito el codigo
func AllowMethods(m ...string) []string {
	return m
}

// Use anida los middleware y se los asigna a un controlador
func Use(handler controller.ControllerFunc, middlewares ...middleware.MiddlewareFunc) controller.ControllerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// HandlerAdapter adapta el controller.Controllerfunc a http.HandlerFunc y envuelve un handler para verificar múltiples métodos HTTP
func HandlerAdapter(handler controller.ControllerFunc, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Si no se especifican métodos, permite cualquiera
		if len(methods) > 0 {

			// Verifica si el método de la solicitud está permitido
			methodAllowed := false
			for _, method := range methods {
				if r.Method == method {
					methodAllowed = true
					break
				}
			}

			if !methodAllowed {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}
		ctx := controller.NewContext(w, r)
		handler(ctx)
	}
}

func HandleFunc(prefix string, routes []Route, middlewares ...middleware.MiddlewareFunc) {

	for _, r := range routes {
		// aplicamos los middlewares
		finalController := Use(r.Handler, middlewares...)

		// adapto el controller para mux
		httpHandler := HandlerAdapter(finalController, r.Methods...)

		// Finalmente registrar en el router
		router.HandleFunc(prefix+r.Path, httpHandler)
	}

}
