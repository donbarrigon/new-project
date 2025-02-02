package routesmaker

import (
	"net/http"

	"github.com/donbarrigon/new-project/internal/controller"
)

// constantes para poder escribir rutas bonitas
const (
	GET     = http.MethodGet     // Solicita la representación de un recurso.
	POST    = http.MethodPost    // Envía datos al servidor para procesarlos.
	PUT     = http.MethodPut     // Reemplaza completamente el recurso con los datos proporcionados.
	DELETE  = http.MethodDelete  // Elimina un recurso.
	PATCH   = http.MethodPatch   // Aplica modificaciones parciales a un recurso.
	HEAD    = http.MethodHead    // Similar a GET, pero solo recupera los encabezados.
	OPTIONS = http.MethodOptions // Describe las opciones de comunicación para el recurso.
	CONNECT = http.MethodConnect // Establece un túnel para comunicaciones (usado principalmente con proxies HTTPS).
	TRACE   = http.MethodTrace   // Realiza una prueba de bucle de retorno para el recurso.
)

// Route estructura de rutas
type Route struct {
	Path    string
	Methods []string
	Handler controller.ControllerFunc
	Name    string
}

// AllowMethods funcion auxiliar para crear un slice de strings y que se vea bonito el codigo
func AllowMethods(m ...string) []string {
	return m
}
