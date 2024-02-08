package main

import (
	"context"
	"simplestack/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")

	return component.Render(context.Background(), c.Response())

}

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return render(c, templates.Home())
	})
	e.Logger.Fatal(e.Start(":8000"))
}
