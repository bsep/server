package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/skratchdot/open-golang/open"
	"strconv"
)

func createAuth(cred []string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
		if username == cred[0] && password == cred[1] {
			return nil, true
		}
		return nil, false
	})
}

func startServer(cfg serverConfig) {
	port := strconv.Itoa(cfg.port)
	e := echo.New()

	if cfg.debug {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	e.GET("/list", handleAll)

	admin := e.Group("/admin")
	admin.Use(createAuth(cfg.credentials))
	admin.GET("/pack", handlePack)
	admin.POST("/add", handleAdd)
	admin.POST("/edit", handleEdit)
	admin.GET("/delete", handleDelete)

	e.Static("/plugins", "plugins")
	e.Static("/", "public")

	if cfg.browser {
		open.Run("http://localhost:" + port + "/")
	}
	e.Start(":" + port)
}
