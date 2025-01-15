package routes

var apiPublic = []route{
	{
		Path:    "/",
		Handler: handlerGet(controller.Example.IndexApi),
		Name:    "index",
	},
}

var apiPrivate = []route{
	{
		Path:    "/",
		Handler: handlerGet(controller.Example.IndexApi),
		Name:    "index",
	},
}
