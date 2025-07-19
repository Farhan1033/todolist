package repository

import (
	"errors"
	"time"
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
	var tasks []model.Task
	err := database.DB.Where("user_id = ?", userId).Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
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
	updates := map[string]interface{}{
		"title":        input.Title,
		"description":  input.Description,
		"is_completed": input.IsCompleted,
		"updated_at":   time.Now(),
	}

	result := database.DB.Model(&model.Task{}).
		Where("id = ? AND user_id = ?", id, input.UserID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task tidak dapat ditemukan atau Anda tidak memiliki akses")
	}

	return nil
}

func (r *taskRepo) DeleteTask(id int) error {
	return database.DB.Delete(&model.Task{}, "id = ?", id).Error
}
