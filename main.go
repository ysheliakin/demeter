package main

import (
	"context"
	"html/template"
	"io"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demeter/db/generated"
	"demeter/handlers"
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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("wwwroot"))
	e.Renderer = newTemplate("views/*/*.html")

	dbc := context.Background()
	conn, err := pgx.Connect(dbc, os.Getenv("DB"))
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	defer conn.Close(dbc)
	query := queries.New(conn)

	e.GET("/post/:id", func(c echo.Context) error {
		return handlers.GetDonationPost(dbc, query, c, e.Logger)
	})
	e.GET("/request/:id", func(c echo.Context) error {
		return handlers.RequestForm(dbc, query, c, e.Logger)
	})
	e.POST("/request/:id", func(c echo.Context) error {
		return handlers.CreateRequest(dbc, query, c, e.Logger)
	})
	e.GET("/location", func(c echo.Context) error {
		return handlers.GetLocationResults(dbc, query, c, e.Logger)
	})
	e.POST("/validate", func(c echo.Context) error {
		return handlers.ValidateFormInput(dbc, query, c, e.Logger)
	})
	e.GET("/sign-in", func(c echo.Context) error {
		return c.Render(200, "sign-in", nil)
	})
	e.GET("/sign-up", func(c echo.Context) error {
		return c.Render(200, "sign-up", nil)
	})
	e.GET("/donate", func(c echo.Context) error {
		return c.Render(200, "donate", nil)
	})
	e.POST("/donate", func(c echo.Context) error {
		return handlers.CreateDonation(dbc, query, c, e.Logger)
	})
	e.GET("/feed", func(c echo.Context) error {
		return handlers.GetDonationPosts(dbc, query, c, e.Logger)
	})
	// NOTE: the more nested routes have to go first to not confuse echo
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
