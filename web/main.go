package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-lottery/bootstrap"
	"go-lottery/web/middleware/identity"
	"go-lottery/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Go抽奖系统", "Sizhe")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
