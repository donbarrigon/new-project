package routes

import "github.com/erespereza/new-project/internal/controller"

var apiPublic = []Route{
	{
		Path:    "/user",
		Methods: AllowMethods(get, post, delete),
		Handler: controller.UserIndex,
		Name:    "user-index",
	},
}

var apiPrivate = []Route{
	{
		Path:    "/user/show",
		Methods: AllowMethods(get),
		Handler: controller.UserShow,
		Name:    "user-show",
	},
}
