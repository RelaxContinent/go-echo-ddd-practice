package application

import (
	"fmt"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/domain"
)

type InitialUseCase interface {
	Invoke(param *InitialParams) (result *InitialResult, err error)
}

type initialUseCase struct {
	repository domain.InitialRepository
}

// NewInitialUseCase
// InitialUseCaseのコンストラクタ
func NewInitialUseCase(repository domain.InitialRepository) InitialUseCase {
	return &initialUseCase{
		repository: repository,
	}
}

// Invoke
// InitialUseCaseの実行メソッド
func (uc *initialUseCase) Invoke(param *InitialParams) (result *InitialResult, err error) {
	// repositoryのメソッド実行
	user, err := uc.repository.GetUserByID(domain.UserID(param.UserID))
	if err != nil {
		fmt.Println("GetUserByID Error:", err)
		return nil, err
	}

	return &InitialResult{
		User: *user,
	}, nil
}
