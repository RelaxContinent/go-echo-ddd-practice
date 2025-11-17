package presentation

import (
	"reflect"
	"testing"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/application"
	"github.com/RelaxContinent/go-echo-ddd-practice/initial/domain"
)

func TestNewInitialResponse(t *testing.T) {
	tests := []struct {
		name     string
		result   *application.InitialResult
		expected string
	}{
		{
			name: "valid user info",
			result: &application.InitialResult{
				User: domain.User{
					ID:       123,
					Name:     "John Doe",
					UserName: "johndoe",
					Email:    "john@example.com",
					Website:  "https://example.com",
				},
			},
			expected: "User Info: ID=123, Name=John Doe, UserName=johndoe, Email=john@example.com, Website=https://example.com",
		},
		{
			name: "empty user info",
			result: &application.InitialResult{
				User: domain.User{
					ID:       0,
					Name:     "",
					UserName: "",
					Email:    "",
					Website:  "",
				},
			},
			expected: "User Info: ID=0, Name=, UserName=, Email=, Website=",
		},
		{
			name: "user with special characters",
			result: &application.InitialResult{
				User: domain.User{
					ID:       456,
					Name:     "Jane O'Brien",
					UserName: "jane_obrien",
					Email:    "jane+test@example.com",
					Website:  "https://example.com/~jane",
				},
			},
			expected: "User Info: ID=456, Name=Jane O'Brien, UserName=jane_obrien, Email=jane+test@example.com, Website=https://example.com/~jane",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := NewInitialResponse(tt.result)

			if !reflect.DeepEqual(response.Message, tt.expected) {
				t.Errorf("expected message %q, got %q", tt.expected, response.Message)
			}
		})
	}
}
