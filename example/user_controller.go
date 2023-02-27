package example

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

//go:generate go-annotate -p "example" -f "user_controller.go"

type UserController struct {
}

func (c UserController) Get() mvc.Result {
	return mvc.Response{
		Object: iris.Map{"userID": "123"},
	}
}

func (c UserController) GetList(ctx iris.Context) mvc.Result {
	folderPath := ctx.URLParam("folderPath")
	return mvc.Response{
		Object: iris.Map{"folderPath": folderPath},
	}
}
