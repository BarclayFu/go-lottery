package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-lottery/services"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceUser    services.UserService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}
