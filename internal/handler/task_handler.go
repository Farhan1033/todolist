package handler

import (
	"net/http"
	"to-do-list/internal/services"
	"to-do-list/model"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	svc services.TaskService
}

func NewTaskHandler(r *gin.RouterGroup, svc services.TaskService) {
	h := &TaskHandler{svc: svc}
	r.POST("/create-task", h.CreateTask)
	r.GET("/task", h.GetTask)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input model.Task
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.svc.CreateTask(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Berhasil membuat task baru"})
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	data, err := h.svc.GetTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    data,
	})
}
