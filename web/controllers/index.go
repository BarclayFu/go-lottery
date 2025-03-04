package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-lottery/comm"
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
func (c *IndexController) Get() {
	c.Ctx.Header("Content-Type", "text/html; charset=utf-8")
	c.Ctx.HTML("<html><body>Welcome to go-lottery!,<a href='/public/index.html'>开始抽奖</a></body></html>")
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

func (c *IndexController) GetLogin() {
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=login")
}

func (c *IndexController) GetLogout() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=logout")
}

func (c *IndexController) GetNewPrize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = "success"
	// TODO
	return rs
}
