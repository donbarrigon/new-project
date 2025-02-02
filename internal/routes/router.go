package routes

import (
	"net/http"

	"github.com/donbarrigon/new-project/internal/middleware"
	"github.com/donbarrigon/new-project/internal/pkg/user"
)

// Create a new router instance
var router = http.NewServeMux()

func NewRouter() *http.ServeMux {

	// rutas para pkg de usuario
	HandleFuncs("/users", user.PublicRoutes())
	HandleFuncs("/users", user.PrivateRoutes(), middleware.Logger, middleware.Request)

	//rutas api standar
	// HandleFuncs("/api/v1", ApiPublic)
	// HandleFuncs("/api/v1", ApiPrivate, middleware.Logger, middleware.Request)

	return router
}
