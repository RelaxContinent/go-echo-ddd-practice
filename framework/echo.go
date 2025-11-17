package framework

import (
	"github.com/labstack/echo/v4"
)

func NewEcho(
	binder echo.Binder,
	validator echo.Validator,
	logger echo.Logger,
) *echo.Echo {
	e := echo.New()
	e.Binder = binder
	e.Validator = validator
	e.Logger = logger
	return e
}
