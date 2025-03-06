package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-lottery/comm"
	"go-lottery/services"
	"go-lottery/web/viewmodels"
)

type AdminGiftController struct {
	Ctx            iris.Context
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceUser    services.UserService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

func (c *AdminGiftController) Get() mvc.Result {
	// TODO:datalist, total
	datalist := c.ServiceGift.GetAll()
	total := len(datalist)
	for i, giftInfo := range datalist {
		// 奖品发放的计划数据
		prizedata := make([][2]int, 0)
		err := json.Unmarshal([]byte(giftInfo.PrizeData), &prizedata)
		if err != nil || len(prizedata) < 1 {
			datalist[i].PrizeData = "[]"
		} else {
			newpd := make([]string, len(prizedata))
			for index, pd := range prizedata {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("[%s] : %d", ct, pd[1])
			}
			str, err := json.Marshal(newpd)
			if err != nil && len(str) > 0 {
				datalist[i].PrizeData = string(str)
			} else {
				datalist[i].PrizeData = "[]"
			}
		}
	}
	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) GetEdit() mvc.Result {
	// TODO: giftInfo
	id := c.Ctx.URLParamIntDefault("id", 0)
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) PostSave() mvc.Result {
	// TODO:
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetDelete() mvc.Result {
	// TODO:
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetReset() mvc.Result {
	// TODO:
	return mvc.Response{
		Path: "/admin/gift",
	}
}
