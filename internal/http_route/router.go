package http_route

import (
	"github.com/mr-chelyshkin/micro/internal/http_route/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpRouter struct {
	e *echo.Echo
}

func NewHttpRouter() *HttpRouter {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", handlers.CurrentTime)

	return &HttpRouter{
		e: e,
	}
}

func (hr HttpRouter) Start() {
	hr.e.Logger.Fatal(
		hr.e.Start(":8090"))
}
