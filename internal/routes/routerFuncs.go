package routes

import (
	"encoding/json"
	"net/http"

	"github.com/donbarrigon/new-project/internal/controller"
	"github.com/donbarrigon/new-project/internal/middleware"
	routesmaker "github.com/donbarrigon/new-project/internal/routes/maker"
)

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
				// se establece el header para la respuesta
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)

				// Crear el mensaje de error en formato JSON
				errorResponse := controller.ErrorResponse{
					Message:    "Method not allowed",
					StatusCode: http.StatusMethodNotAllowed,
				}

				// Serializar el objeto errorResponse a JSON
				json.NewEncoder(w).Encode(errorResponse)
				return
			}
		}
		ctx := controller.NewContext(w, r)
		handler(ctx)
	}
}

func HandleFuncs(prefix string, routes []routesmaker.Route, middlewares ...middleware.MiddlewareFunc) {

	for _, r := range routes {
		// aplicamos los middlewares
		finalController := Use(r.Handler, middlewares...)

		// adapto el controller para mux
		httpHandler := HandlerAdapter(finalController, r.Methods...)

		// Finalmente registrar en el router
		router.HandleFunc(prefix+r.Path, httpHandler)
	}

}
