package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", HelloHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}
