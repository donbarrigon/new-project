package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func UserIndex(ctx *Context) {
	// user := model.NewUser()
	// user.Find(1)
	response := Response{
		Message: "Request was successful desde index",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func UserShow(ctx *Context) {
	// user := model.NewUser()
	// user.Find(1)
	response := Response{
		Message: "Request was successful desde show",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}
