package main

import (
	"to-do-list/config"
	"to-do-list/database"
	"to-do-list/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":" + config.Get("PORT"))
}
