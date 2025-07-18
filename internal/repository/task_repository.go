package repository

import (
	"errors"
	"to-do-list/database"
	"to-do-list/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(*model.Task) error
	GetTask() ([]model.Task, error)
	GetTaskByUserId(userId uuid.UUID) ([]model.Task, error)
	GetTaskById(id int) (*model.Task, error)
	UpdateTask(id int, input *model.Task) error
	DeleteTask(id int) error
}

type taskRepo struct{}

func NewTaskRepo() TaskRepository {
	return &taskRepo{}
}

func (r *taskRepo) CreateTask(task *model.Task) error {
	return database.DB.Create(task).Error
}

func (r *taskRepo) GetTask() ([]model.Task, error) {
	var task []model.Task
	err := database.DB.Find(&task).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return task, nil
}

func (r *taskRepo) GetTaskByUserId(userId uuid.UUID) ([]model.Task, error) {
	var task []model.Task
	err := database.DB.Where("user_id = ?", userId).Find(&task).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return task, nil
}

func (r *taskRepo) GetTaskById(id int) (*model.Task, error) {
	var task model.Task
	err := database.DB.First(&task, "id = ?", id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &task, nil
}

func (r *taskRepo) UpdateTask(id int, input *model.Task) error {
	var task model.Task
	err := database.DB.First(&task, "id = ?", id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	task.ID = input.ID
	task.Title = input.Title
	task.Description = input.Description
	task.IsCompleted = input.IsCompleted
	task.CreatedAt = input.CreatedAt
	task.UpdatedAt = input.UpdatedAt

	return database.DB.Save(&task).Error
}

func (r *taskRepo) DeleteTask(id int) error {
	return database.DB.Delete(&model.Task{}, "id = ?", id).Error
}
