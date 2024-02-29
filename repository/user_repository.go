package repository

import "intikom_test/entity"

type (
	UserRepository interface {
		Create(model entity.User) (uint, error)
		Get(id string) (*entity.User, error)
		List(limit, offset int) ([]entity.User, int64, error)
		Update(model *entity.User) error
		Delete(model *entity.User) error
	}
)
