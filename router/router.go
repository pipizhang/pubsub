package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"

	ctl "github.com/pipizhang/pubsub/controller"
	mw "github.com/pipizhang/pubsub/middleware"
	"github.com/pipizhang/pubsub/pkg"
)

func Start(confFile string, ipDBFile string) {

	pkg.InitConf(confFile)
	pkg.InitLog()

	e := echo.New()

	e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${remote_ip} \"${method} ${host}${path}\" - ${status} ${latency_human}\n",
	}))
	e.Pre(mw.ServerHeader())
	e.Pre(mw.IPWhitelistWithConfig(mw.IPWhitelistConfig{
		Enable:   pkg.Conf.IPWhitelist.Enable,
		List:     pkg.Conf.IPWhitelist.List,
		IPDBFile: ipDBFile,
	}))
	e.Use(middleware.Recover())

	e.GET("/api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/api/publish", ctl.BrokerController.Publish)

	e.Server.Addr = pkg.Conf.Server.Address

	fmt.Println(fmt.Sprintf("Start %s ...", pkg.Conf.App.Name))
	pkg.AppLog.Info(fmt.Sprintf("Start %s ...", pkg.Conf.App.Name))
	defer func() {
		pkg.AppLog.Info("Shutdown")
	}()

	err := graceful.ListenAndServe(e.Server, pkg.Conf.ServerOffTimeout())
	if err != nil {
		pkg.AppLog.Fatal(err)
	}
}
