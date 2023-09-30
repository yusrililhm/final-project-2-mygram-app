package handler

import (
	"myGram/infra/config"
	"myGram/infra/database"

	"github.com/gin-gonic/gin"
)

func StartApplication() {

	config.LoadEnv()

	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()
	_ = db

	app := gin.Default()

	app.Run(":" + config.AppConfig().Port)
}
