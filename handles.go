package main

import (
	"simplego/views"

	"github.com/labstack/echo"
)

func ShowHomePage(c echo.Context) error {
	component := views.CompaniesView()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
