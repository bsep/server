package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func createAuth(cred []string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
		if username == cred[0] && password == cred[1] {
			return nil, true
		}
		return nil, false
	})
}

func server(port string, cred []string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/list", handleAll)

	admin := e.Group("/admin")
	admin.Use(createAuth(cred))
	admin.GET("/pack", handlePack)
	admin.POST("/add", handleAdd)
	admin.POST("/edit", handleEdit)
	admin.GET("/delete", handleDelete)

	e.Static("/plugins", "plugins")
	e.Static("/", "public")

	e.Start(":" + port)
}

var version = "master"

func main() {
	fmt.Println("Search Plugins", version, "by Doğan Çelik (dogancelik.com)")
	fmt.Println("Starting Search Plugins server…")

	portPtr := flag.Int("p", 8080, "server port")
	credPtr := flag.String("c", "admin:123456", "admin username & password")
	flag.Parse()

	port := fmt.Sprintf("%d", *portPtr)
	cred := *credPtr

	errCred, credArr := splitCreds(cred)
	if errCred != nil {
		fmt.Println(fmt.Errorf("Error: %v", errCred))
		os.Exit(1)
	}

	server(port, credArr)
}
