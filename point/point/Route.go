package point

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	point := &Point{}
	e.GET("/points", point.Get)
	e.GET("/points/:id", point.GetbyId)
	e.POST("/points", point.Persist)
	e.PUT("/points/:id", point.Put)
	e.DELETE("/points/:id", point.Remove)
	e.PUT("/points/use", Use)
	return e
}
