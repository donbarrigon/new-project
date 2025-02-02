package user

import (
	. "github.com/donbarrigon/new-project/internal/routes/maker"
)

func PublicRoutes() []Route {
	return []Route{
		{
			Path:    "/",
			Methods: AllowMethods(GET),
			Handler: IndexController,
			Name:    "user-index",
		},
	}
}

func PrivateRoutes() []Route {
	return []Route{
		{
			Path:    "/show",
			Methods: AllowMethods(GET),
			Handler: ShowController,
			Name:    "user-show",
		},
		{
			Path:    "/create",
			Methods: AllowMethods(POST),
			Handler: CreateController,
			Name:    "user-create",
		},
		{
			Path:    "/update",
			Methods: AllowMethods(PUT),
			Handler: UpdateController,
			Name:    "user-update",
		},
		{
			Path:    "/delete",
			Methods: AllowMethods(DELETE),
			Handler: DeleteController,
			Name:    "user-delete",
		},
	}
}
