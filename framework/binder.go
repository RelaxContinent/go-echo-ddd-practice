package framework

import "github.com/labstack/echo/v4"

type CustomBindable interface {
	Bind(c echo.Context) error
}

type CustomBinder struct {
	echo.Binder
}

func NewCustomBinder() echo.Binder {
	return &CustomBinder{
		Binder: &echo.DefaultBinder{},
	}
}
