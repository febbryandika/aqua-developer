package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type Name struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	log.SetFormatter(&log.JSONFormatter{})
	log.Printf("Run Echo")
	e.GET("/users/:name", func(c echo.Context) error {
		name := &User{
			Success: true,
			Data: Name{
				Name: c.Param("name"),
			},
		}
		return c.JSON(http.StatusOK, name)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
