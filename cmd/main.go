package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/http"
	"strconv"
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

	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	c := rc.Ping(context.Background())
	if c.Err() != nil {
		panic(c.Err())
	}

	v1.GET("/user", func(c echo.Context) error {
		rc.Set(context.Background(), strconv.Itoa(rand.Int()), "user", time.Duration(time.Minute))
		return c.JSON(http.StatusOK, map[string]string{
			"url": "user",
			"ok":  "true",
		})
	})
	v1.GET("/profile", func(c echo.Context) error {
		res := rc.Keys(context.Background(), "*")
		strs, err := res.Result()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		m := make(map[string]string)

		for _, s := range strs {
			val := rc.Get(context.Background(), s)

			m[s] = val.String()
		}

		return c.JSON(http.StatusOK, m)
	})

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
