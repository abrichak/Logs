package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func SuccessResponse(c echo.Context, message string) error {
	return Response(c, http.StatusOK, Success{
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
