package handler

import (
	"myGram/infra/config"
	"myGram/infra/database"
	"myGram/repository/user_repository/user_pg"
	"myGram/service/user_service"

	"github.com/gin-gonic/gin"
)

func StartApplication() {

	config.LoadEnv()

	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()

	// dependencies injection
	userRepo := user_pg.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	app := gin.Default()

	// swagger

	// routing
	users := app.Group("users")

	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
	}

	app.Run(":" + config.AppConfig().Port)
}
