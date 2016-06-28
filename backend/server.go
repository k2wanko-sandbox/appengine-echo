package backend

import (
	"net/http"

	ae "github.com/k2wanko/echo-appengine"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"google.golang.org/appengine"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(ae.AppContext())
	e.Use(ae.AppLogger())
	e.Use(Logger())
	e.GET("/", handleIndex)
	return e
}

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	s := standard.New("")
	s.SetHandler(New())
	http.Handle("/", s)
}

func handleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "index: "+appengine.AppID(c))
}

func AppContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if r, ok := c.Request().(*standard.Request); ok {
				c.SetContext(appengine.WithContext(c.Context(), r.Request))
			}
			return next(c)
		}
	}
}
