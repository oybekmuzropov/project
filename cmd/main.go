package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	defer func() {
		recover()
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	v1 := e.Group("/api/v1")

	v1.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"url": "user",
			"ok":  "true",
		})
	})
	v1.GET("/profile", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"url": "profile",
			"ok":  "true",
		})
	})

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
