package backend

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"google.golang.org/appengine/log"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req, res := c.Request(), c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			end := time.Now()
			latency := end.Sub(start)
			ip := req.RemoteAddress()
			method := req.Method()
			p := req.URL().Path()
			if p == "" {
				p = "/"
			}
			s := res.Status()
			o := fmt.Sprintf("%v | %3d | %s | %s | %-7s | %s",
				end.Format("2006/01/02 - 15:04:05"),
				s,
				latency.String(),
				ip,
				method,
				p,
			)
			switch {
			case s >= 500:
				log.Errorf(c, "%s", o)
			case s >= 400:
				log.Warningf(c, "%s", o)
			default:
				log.Infof(c, "%s", o)
			}
			return
		}
	}
}
