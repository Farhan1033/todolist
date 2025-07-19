package handler

import (
	"net/http"
	"strconv"
	"strings"
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
	r.GET("/task-user", h.GetTaskUser)
	r.GET("/task/:id", h.GetTaskId)
	r.PUT("/update-task/:id", h.UpdateTask)
	r.DELETE("/delete-task/:id", h.DeleteTask)
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

func (h *TaskHandler) GetTaskUser(c *gin.Context) {
	data, err := h.svc.GetTaskByUserId(c)
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

func (h *TaskHandler) GetTaskId(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak sesuai"})
		return
	}

	data, err := h.svc.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error saat ambil data"})
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

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input model.Task
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data bukan JSON yang valid"})
		return
	}

	err = h.svc.UpdateTask(id, &input, c)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "tidak diizinkan") || strings.Contains(err.Error(), "unauthorized") {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil update data"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	delete := h.svc.DeleteTask(id)
	if delete != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": delete.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil hapus data"})
}
