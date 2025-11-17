package main

import (
	"github.com/RelaxContinent/go-echo-ddd-practice/framework"
	"github.com/RelaxContinent/go-echo-ddd-practice/initial"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		initial.Modules,
		framework.Modules,
		fx.Invoke(func(e *echo.Echo) {
			e.Logger.Fatal(e.Start(":1323"))
		}),
	)
	app.Run()
}
