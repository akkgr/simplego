package main

import (
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/", ShowHomePage)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
