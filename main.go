package main

import (
	"bjj_cms/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	h := handlers.New()
	e := echo.New()
	e.File("/style.css", "./static/style.css")
	e.GET("/", h.ViewCollections)
	e.Logger.Fatal(e.Start(":1223"))
}
