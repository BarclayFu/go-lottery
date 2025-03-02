package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"go-lottery/bootstrap"
	"go-lottery/services"
	"go-lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewblackipService()

	index := mvc.New(b.Party("/"))
	index.Register(userService, userdayService, giftService, codeService, resultService, blackipService)
	index.Handle(new(controllers.IndexController))
}
