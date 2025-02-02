package user

import (
	"encoding/json"
	"net/http"

	"github.com/donbarrigon/new-project/internal/controller"
)

type Response struct {
	Message string `json:"message"`
}

func IndexController(ctx *controller.Context) {

	response := Response{
		Message: "IndexController is ready",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func ShowController(ctx *controller.Context) {

	response := Response{
		Message: "ShowController is ready",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func CreateController(ctx *controller.Context) {

	response := Response{
		Message: "CreateController is ready",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateController(ctx *controller.Context) {

	response := Response{
		Message: "UpdateController is ready",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteController(ctx *controller.Context) {

	response := Response{
		Message: "DeleteController is ready",
	}
	if err := json.NewEncoder(ctx.Writer).Encode(response); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}
