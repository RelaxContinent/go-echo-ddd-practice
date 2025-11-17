package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/domain"
)

type InitialRepositoryImpl struct{}

func NewInitialRepositoryImpl() domain.InitialRepository {
	return &InitialRepositoryImpl{}
}

// GetUserByID
// 外部APIからUser情報を取得する
func (r *InitialRepositoryImpl) GetUserByID(id domain.UserID) (*domain.User, error) {
	apiResponse, err := request(id)
	if err != nil {
		fmt.Println("Error occurred while requesting external API:", err)
		return nil, err
	}
	user := mapToDomainObject(apiResponse)
	return user, nil
}

// request
// 外部APIにRequestを送り、Responseを受け取る
func request(id domain.UserID) (*ApiResponse, error) {
	// HTTPクライアントの生成
	client := &http.Client{Timeout: 5 * time.Second}

	// Requestの生成
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/"+string(id), nil)
	if err != nil {
		fmt.Println("Error occurred while creating HTTP request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// APIにRequest送信、Response取得
	apiResponse, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occurred while sending HTTP request:", err)
		return nil, err
	}

	// この関数の終了時に、Response Bodyのstream読み取りの終了を実行する
	defer apiResponse.Body.Close()

	// APIのResponseのStatusCodeが200以外の場合はエラーを返却
	if apiResponse.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK HTTP status:", apiResponse.StatusCode)
		return nil, err
	}

	// Response Bodyを構造体にデコード
	var response ApiResponse
	err = json.NewDecoder(apiResponse.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error occurred while decoding API response:", err)
		return nil, err
	}

	return &response, nil
}

// mapToDomainObject
// APIのResponseをDomainオブジェクトにマッピングする
func mapToDomainObject(apiResponse *ApiResponse) *domain.User {
	return &domain.User{
		ID:       domain.UserID(apiResponse.ID),
		Name:     apiResponse.Name,
		UserName: apiResponse.UserName,
		Email:    apiResponse.Email,
		Website:  apiResponse.Website,
	}
}
