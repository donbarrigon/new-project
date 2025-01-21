package controller

import (
	"net/http"

	"github.com/erespereza/new-project/internal/model"
)

type ControllerFunc func(ctx *Context)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  map[string]string
	User    *model.User
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		// Params:   mux.Vars(r),
	}
}
