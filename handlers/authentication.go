package handlers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

/* Excess Code
store := sessions.NewCookieStore([]byte("secret"))

// User Authentication
//e.Use(session.Middleware(gothic.Store))
store.Options = &sessions.Options{
	HttpOnly: true, // Set cookies as HTTP only
}

// Customized session middleware
e.Use(session.Middleware(store))
*/

func GothUserProvider(port string) {
	goth.UseProviders(
		google.New(os.Getenv("KEY"), os.Getenv("SECRET"), "http://localhost:"+port+"/auth/google/callback"),
	)
}

func GothAuth(c echo.Context) error {
	// Try to get user without re-authenticating
	provider := "google"
	req := gothic.GetContextWithProvider(c.Request(), provider)
	if gothUser, err := gothic.CompleteUserAuth(c.Response(), req); err == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	// Begin authentication with Google provider
	gothic.BeginAuthHandler(c.Response(), req)
	return nil
}

func GothCallback(c echo.Context) error {
	// Complete user authentication
	provider := "google"
	req := gothic.GetContextWithProvider(c.Request(), provider)
	user, err := gothic.CompleteUserAuth(c.Response(), req)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func GothLogout(c echo.Context) error {
	// Logout user from Google provider
	provider := "google"
	req := gothic.GetContextWithProvider(c.Request(), provider)
	gothic.Logout(c.Response(), req)

	// Redirect to home
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
