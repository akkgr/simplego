package main

import (
	"log/slog"

	"github.com/labstack/echo"
)

func main() {
	logger := slog.Default()

	logger.Info("Starting server")

	e := echo.New()

	e.GET("/", ShowHomePage)

	e.Static("/css", "static/css")
	e.Static("/js", "static/js")
	e.Static("/font", "static/font")
	e.Static("/favicon.ico", "static/favicon.ico")

	e.Logger.Fatal(e.Start(":3000"))
}
