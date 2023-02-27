package example

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
}

// Get
// @Summary TODO
// @Description TODO
// @Tags TODO
// @Accept json
// @Produce json
// @Param usercontroller path string true "UserController ID"
// @Success 200 {object} &{mvc Result}
// @Router /usercontroller/ [get]
func (c UserController) Get() mvc.Result {
	return mvc.Response{
		Object: iris.Map{"userID": "123"},
	}
}

// GetList
// @Summary TODO
// @Description TODO
// @Tags TODO
// @Accept json
// @Produce json
// @Param usercontroller path string true "UserController ID"
// @Param ctx body &{iris true "ctx"
// @Param ctx query  true "ctx"
// @Success 200 {object} &{mvc Result}
// @Router /usercontroller/List [get]
func (c UserController) GetList(ctx iris.Context) mvc.Result {
	folderPath := ctx.URLParam("folderPath")
	return mvc.Response{
		Object: iris.Map{"folderPath": folderPath},
	}
}
