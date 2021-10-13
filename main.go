package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld() string {
	return "hello world"
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%s", HelloWorld()))
}

func initService() {
	e := echo.New()

	e.GET("/hello", hello)

	e.Start(":8005")
}

func main() {
	initService()
}
