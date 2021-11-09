package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func CurrentTime(c echo.Context) error {
	type response struct {
		Time time.Time `json:"time"`
	}
	resp := response{
		Time: time.Now(),
	}

	return c.JSON(http.StatusOK, resp)
}
