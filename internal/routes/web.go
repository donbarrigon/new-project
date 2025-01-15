package routes

var webPublic = []route{
	{
		Path:    "/",
		Handler: handlerGet(Index),
		Name:    "index",
	},
}

var webPrivate = []route{
	{
		Path:    "/",
		Handler: handlerGet(Index),
		Name:    "index",
	},
}
