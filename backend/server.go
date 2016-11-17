package backend

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(Logger())
	e.GET("/", handleIndex)

	http.Handle("/", e)
}

func handleIndex(c echo.Context) error {
	ctx := newContext(c)
	return c.String(http.StatusOK, "index: "+appengine.AppID(ctx))
}

func newContext(c echo.Context) context.Context {
	return appengine.NewContext(c.Request())
}
