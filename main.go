package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	Templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func newTemplate(path string) *TemplateRegistry {
	return &TemplateRegistry{
		Templates: template.Must(template.ParseGlob(path)),
	}
}

type Payload struct {
	Message string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate("views/*.html")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", Payload{Message: "Hello World"})
	})

	e.Logger.Fatal(e.Start(":42069"))
}
