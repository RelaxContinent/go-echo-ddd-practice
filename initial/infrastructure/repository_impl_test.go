package infrastructure

import (
	"reflect"
	"testing"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/domain"
)

func Test_mapToDomainObject(t *testing.T) {
	tests := []struct {
		name        string
		apiResponse *ApiResponse
		want        *domain.User
	}{
		{
			name: "正常系: ApiResponseをdomain.Userにマッピングできること",
			apiResponse: &ApiResponse{
				ID:       1,
				Name:     "Leanne Graham",
				UserName: "Bret",
				Email:    "Sincere@april.biz",
				Website:  "hildegard.org",
			},
			want: &domain.User{
				ID:       1,
				Name:     "Leanne Graham",
				UserName: "Bret",
				Email:    "Sincere@april.biz",
				Website:  "hildegard.org",
			},
		},
		{
			name:        "正常系: ApiResponseがnilの場合、domain.Userもnilを返すこと",
			apiResponse: nil,
			want:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapToDomainObject(tt.apiResponse)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToDomainObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
