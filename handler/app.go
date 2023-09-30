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
	userHandler := newUserHandler(userService)

	app := gin.Default()

	// routing
	users := app.Group("users")
	
	{
		users.POST("/register", userHandler.Register)
	}

	app.Run(":" + config.AppConfig().Port)
}
