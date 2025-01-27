package middleware

import (
	"fmt"

	"github.com/donbarrigon/new-project/internal/controller"
)

type MiddlewareFunc func(controller.ControllerFunc) controller.ControllerFunc

func Logger(next controller.ControllerFunc) controller.ControllerFunc {
	return func(ctx *controller.Context) {
		fmt.Printf("Before %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
		fmt.Printf("After %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
	}
}

func Request(next controller.ControllerFunc) controller.ControllerFunc {
	return func(ctx *controller.Context) {
		fmt.Printf("Request: %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
	}
}
