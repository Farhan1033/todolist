package router

import (
	"to-do-list/internal/handler"
	"to-do-list/internal/repository"
	"to-do-list/internal/services"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	userRepo := repository.NewUserRepo()
	userSvc := services.NewUserService(userRepo)

	api := r.Group("/api/v1")
	{
		handler.NewUserHandler(api, userSvc)
	}
}
