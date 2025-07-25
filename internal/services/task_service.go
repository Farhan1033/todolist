package services

import (
	"errors"
	"to-do-list/internal/repository"
	"to-do-list/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(*gin.Context, *model.Task) error
	GetTask() ([]model.Task, error)
	GetTaskByUserId(c *gin.Context) ([]model.Task, error)
	GetTaskById(id int) (*model.Task, error)
	UpdateTask(id int, input *model.Task, c *gin.Context) error
	DeleteTask(id int) error
}

type taskSvc struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskSvc{repo: r}
}

func (s *taskSvc) CreateTask(c *gin.Context, task *model.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}

	if task.Description == "" {
		return errors.New("description cannot be empty")
	}

	userId, err := c.Value("user_id").(uuid.UUID)
	if !err {
		userIDStr, exists := c.Value("user_id").(string)
		if !exists {
			return errors.New("user not authenticated")
		}

		parsedUserID, err := uuid.Parse(userIDStr)
		if err != nil {
			return errors.New("invalid user ID format")
		}
		userId = parsedUserID
	}

	task.UserID = userId
	return s.repo.CreateTask(task)
}

func (s *taskSvc) GetTask() ([]model.Task, error) {
	datas, err := s.repo.GetTask()

	if err != nil {
		return nil, err
	}

	if len(datas) == 0 {
		return nil, errors.New("task not found")
	}

	return datas, err
}

func (s *taskSvc) GetTaskByUserId(c *gin.Context) ([]model.Task, error) {
	userId, ok := c.Value("user_id").(uuid.UUID)
	if !ok {
		userIdStr, exists := c.Value("user_id").(string)
		if !exists {
			return nil, errors.New("user tidak ditemukan")
		}

		parsedId, err := uuid.Parse(userIdStr)
		if err != nil {
			return nil, errors.New("token tidak valid")
		}
		userId = parsedId
	}

	tasks, err := s.repo.GetTaskByUserId(userId)
	if err != nil {
		return nil, errors.New("gagal mengambil data task")
	}

	if len(tasks) == 0 {
		return nil, errors.New("task not found")
	}

	return tasks, nil
}

func (s *taskSvc) GetTaskById(id int) (*model.Task, error) {
	tasks, err := s.repo.GetTaskById(id)

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		return nil, errors.New("task not found")
	}
	return tasks, err
}

func (s *taskSvc) UpdateTask(id int, input *model.Task, c *gin.Context) error {
	if input.Title == "" || input.Description == "" {
		return errors.New("fields cannot be empty")
	}

	userId, ok := c.Value("user_id").(uuid.UUID)
	if !ok {
		userIdStr, exists := c.Value("user_id").(string)
		if !exists {
			return errors.New("user tidak ditemukan")
		}

		parsedId, err := uuid.Parse(userIdStr)
		if err != nil {
			return errors.New("token tidak valid")
		}
		userId = parsedId
	}

	input.UserID = userId

	return s.repo.UpdateTask(id, input)
}

func (s *taskSvc) DeleteTask(id int) error {
	if id == 0 {
		return errors.New("id cannot be empty")
	}
	return s.repo.DeleteTask(id)
}
