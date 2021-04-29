package presenter

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		uh      *UserHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uh.GetAllUsers(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("UserHandler.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
