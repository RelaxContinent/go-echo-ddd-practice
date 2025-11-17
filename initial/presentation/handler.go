package presentation

import (
	"fmt"
	"net/http"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/application"
	"github.com/labstack/echo/v4"
)

type InitialHandler struct {
	initialUseCase application.InitialUseCase
}

func NewInitialHandler(initialUseCase application.InitialUseCase) *InitialHandler {
	return &InitialHandler{
		initialUseCase: initialUseCase,
	}
}

// GetUser
// ユーザ情報取得APIのハンドラ
func (h *InitialHandler) GetUser(c echo.Context) error {
	// Request取得
	request := new(InitialRequest)
	if err := c.Bind(request); err != nil {
		fmt.Println("Bind Error:", err)
		return err
	}
	if err := c.Validate(request); err != nil {
		fmt.Println("Validate Error:", err)
		return err
	}
	// Application層のパラメータに変換
	params := request.ToApplicationParams()
	result, err := h.initialUseCase.Invoke(params)
	if err != nil {
		fmt.Println("Invoke Error:", err)
		return err
	}
	// Response返却
	response := NewInitialResponse(result)
	if err := c.JSON(http.StatusOK, response); err != nil {
		fmt.Println("JSON Response Error:", err)
		return err
	}
	return nil
}
