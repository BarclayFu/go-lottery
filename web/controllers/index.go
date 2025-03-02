package controllers

import (
	"github.com/kataras/iris/v12"
	"go-lottery/models"
	services "go-lottery/services"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceUser    services.UserService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// http://localhost:8080
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "Welcome to go-lottery!,<a herf='/public/index.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = "success"
	datalist := c.ServiceGift.GetAll()
	list := make([]models.LtGift, 0)

	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}

	rs["gift"] = list
	return rs
}

func (c *IndexController) GetNewPrize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = "success"
	// TODO
	return rs
}
