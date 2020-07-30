package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"makeToon/api"
	"makeToon/config"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: config.ErrorFormat() + "\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string {"*"},
		AllowHeaders: []string {echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		AllowMethods: []string {echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	v1 := e.Group("/api/v1")
	{
		v1.PUT("/photo", api.PutCropPhoto)
		v1.GET("/photo", api.GetMapPhotos)
	}

	return e
}
