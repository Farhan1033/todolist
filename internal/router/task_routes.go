package router

import (
	"to-do-list/internal/handler"
	"to-do-list/internal/middleware"
	"to-do-list/internal/repository"
	"to-do-list/internal/services"

	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(r *gin.Engine){
	taskRepo := repository.NewTaskRepo()
	taskSvc := services.NewTaskService(taskRepo)

	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth())
	{
		handler.NewTaskHandler(api, taskSvc)
	}
}