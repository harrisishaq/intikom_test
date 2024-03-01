package repository

import (
	"errors"
	"intikom_test/entity"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	taskRepository struct {
		DB *gorm.DB
	}
)

func NewTaskRepository(conn *gorm.DB) TaskRepository {
	conn.Model(&entity.Task{})

	return &taskRepository{
		DB: conn,
	}
}

func (repo *taskRepository) Create(model entity.Task) (uint, error) {
	db := repo.DB.Create(&model)
	err := db.Error
	if err != nil {
		log.Println("error cause: ", err)
		return 0, err
	}
	return model.ID, nil
}

func (repo *taskRepository) Delete(model *entity.Task) error {
	qResult := repo.DB.Clauses(clause.Returning{}).Delete(model)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil
	} else if qResult.Error != nil {
		return qResult.Error
	}
	return nil
}

func (repo *taskRepository) Get(id string) (*entity.Task, error) {
	var result entity.Task
	qResult := repo.DB.First(&result, "id = ?", id)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if qResult.Error != nil {
		return nil, qResult.Error
	}
	return &result, nil
}

func (repo *taskRepository) List(limit, offset int) ([]entity.Task, int64, error) {
	var results []entity.Task
	var tx = repo.DB.Model(&entity.Task{})
	tx.Order("id asc")

	if limit > 0 {
		tx.Limit(limit)
		tx.Offset(offset)
	}

	err := tx.Find(&results).Error
	if err != nil {
		return make([]entity.Task, 0), 0, err
	}

	var total int64
	err = tx.Count(&total).Error
	if err != nil {
		return make([]entity.Task, 0), 0, err
	}

	return results, total, nil
}

func (repo *taskRepository) Update(model *entity.Task) error {
	qResult := repo.DB.Select("*").Where("id = ?", model.ID).Updates(model)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil
	} else if qResult.Error != nil {
		return qResult.Error
	}
	return nil
}
