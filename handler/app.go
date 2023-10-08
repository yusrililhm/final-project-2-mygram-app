package handler

import (
	"myGram/infra/config"
	"myGram/infra/database"
	"myGram/repository/comment_repository/comment_pg"
	"myGram/repository/photo_repository/photo_pg"
	"myGram/repository/user_repository/user_pg"
	"myGram/service/auth_service"
	"myGram/service/comment_service"
	"myGram/service/photo_service"
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

	photoRepo := photo_pg.NewPhotoRepository(db)
	photoService := photo_service.NewPhotoService(photoRepo)
	photoHandler := NewPhotoHandler(photoService)

	commentRepo := comment_pg.NewCommentRepository(db)
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)
	commentHandler := NewCommentHandler(commentService)

	authService := auth_service.NewAuthService(userRepo, photoRepo, commentRepo)

	app := gin.Default()

	// swagger

	// routing
	users := app.Group("users")

	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
		users.PUT("", authService.Authentication(), userHandler.Update)
		users.DELETE("", authService.Authentication(), userHandler.Delete)
	}

	photos := app.Group("photos")

	{
		photos.POST("", authService.Authentication(), photoHandler.AddPhoto)
		photos.GET("", authService.Authentication(), photoHandler.GetPhotos)
		photos.PUT("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.UpdatePhoto)
		photos.DELETE("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.DeletePhoto)
	}

	comments := app.Group("comments")

	{
		comments.POST("", authService.Authentication(), commentHandler.AddComment)
		comments.GET("", authService.Authentication(), commentHandler.GetComments)
		comments.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateComment)
		comments.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteComment)
	}

	app.Run(":" + config.AppConfig().Port)
}
