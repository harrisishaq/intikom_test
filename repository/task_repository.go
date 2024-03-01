package repository

import "intikom_test/entity"

type (
	TaskRepository interface {
		Create(model entity.Task) (uint, error)
		Delete(model *entity.Task) error
		Get(id string) (*entity.Task, error)
		List(limit, offset int) ([]entity.Task, int64, error)
		Update(model *entity.Task) error
	}
)
