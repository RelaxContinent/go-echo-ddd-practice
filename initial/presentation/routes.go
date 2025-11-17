package presentation

import "github.com/labstack/echo/v4"

// SetUpRoute
// ルーティング設定
func SetUpRoute(e *echo.Echo, handler *InitialHandler) {
	e.GET("/initial", handler.GetUser)
}
