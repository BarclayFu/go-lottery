package bootstrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"go-lottery/conf"
	"time"
)

type Configurator func(bootstrapper *Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

func (b *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true)
	htmlEngine.AddFunc("FromUnixtimeShort", func(timestamp int) string {
		dt := time.Unix(int64(timestamp), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(timestamp int) string {
		dt := time.Unix(int64(timestamp), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}
		if jsonPutput := ctx.URLParamExists("json"); jsonPutput {
			ctx.JSON(err)
			return
		}
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) setupCron() {
	//TODO: 计划任务
}

const (
	StaticAssets = "./public/"
	Favicon      = "favicon.ico"
)

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupErrorHandlers()

	b.Favicon(StaticAssets + Favicon)
	b.HandleDir(StaticAssets[1:len(StaticAssets)], StaticAssets)

	b.setupCron()

	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
