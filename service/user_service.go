package service

import "intikom_test/model"

type (
	UserService interface {
		CreateUser(req *model.CreateUserRequest) error
		DeleteUser(id string) error
		GetUser(id string) (*model.DataUserResponse, error)
		ListUser() ([]model.DataUserResponse, int64, error)
		UpdateUser(req *model.UpdateUserRequest) error
	}
)
