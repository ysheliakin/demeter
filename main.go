package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demeter/db/generated"
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
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("wwwroot"))

	dbc := context.Background()
	conn, err := pgx.Connect(dbc, os.Getenv("DB"))
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	defer conn.Close(dbc)
	query := queries.New(conn)

	user, err := query.GetUser(dbc, 1)
	if err != nil {
		e.Logger.Errorf("failed to retrieve user: %s\n", err)
	}
	fmt.Println(user)

	e.Renderer = newTemplate("views/*.html")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", Payload{Message: "Hello World"})
	})
	e.GET("/ui", func(c echo.Context) error {
		return c.Render(200, "ui", nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
