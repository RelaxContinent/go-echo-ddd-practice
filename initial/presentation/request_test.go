package presentation_test

import (
	"reflect"
	"testing"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/application"
	"github.com/RelaxContinent/go-echo-ddd-practice/initial/presentation"
)

func TestInitialRequest_ToApplicationParams(t *testing.T) {
	tests := []struct {
		name     string
		receiver presentation.InitialRequest
		want     *application.InitialParams
	}{
		{
			name: "userIDが正しく変換されること",
			receiver: presentation.InitialRequest{
				UserID: 5,
			},
			want: &application.InitialParams{
				UserID: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.receiver.ToApplicationParams()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToApplicationParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
