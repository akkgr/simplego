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

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
