package usecase

import (
	"reflect"
	"svc-users-go/module/v1/user/model"
	"svc-users-go/module/v1/user/repository"
	"testing"
)

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		UserRepos repository.Repository
	}
	tests := []struct {
		name string
		args args
		want UseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUseCase(tt.args.UserRepos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCaseImpl_CountAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CountAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.CountAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUseCaseImpl.CountAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCaseImpl_FindAllUsers(t *testing.T) {
	type args struct {
		name  string
		limit int64
		page  int64
	}
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		args    args
		want    []model.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.FindAllUsers(tt.args.name, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.FindAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCaseImpl.FindAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCaseImpl_FindUserById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		args    args
		want    model.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.FindUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.FindUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCaseImpl.FindUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCaseImpl_CreateNewUser(t *testing.T) {
	type args struct {
		payload *model.CreateUser
	}
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		args    args
		want    *model.CreateUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CreateNewUser(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCaseImpl.CreateNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCaseImpl_UpdateUser(t *testing.T) {
	type args struct {
		id      string
		payload *model.UpdateUser
	}
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.UpdateUser(tt.args.id, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUseCaseImpl_DeleteUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		uc      *userUseCaseImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userUseCaseImpl.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
