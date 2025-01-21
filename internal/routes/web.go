package routes

import "github.com/erespereza/new-project/internal/controller"

var webPublic = []Route{
	{
		Path:    "/user",
		Methods: AllowMethods(get),
		Handler: controller.UserIndex,
		Name:    "user-index",
	},
}

var webPrivate = []Route{
	{
		Path:    "/user/show",
		Methods: AllowMethods(get),
		Handler: controller.UserShow,
		Name:    "user-show",
	},
}
