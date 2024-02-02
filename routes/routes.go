package routes

import (
	"github.com/labstack/echo/v4"
	"gormTest/database"
	dbImage "gormTest/database/image"
	dbTag "gormTest/database/tag"
	dbUniverse "gormTest/database/universe"
	handlerImage "gormTest/handlers/image"
	handlerTag "gormTest/handlers/tag"
	handlerUniverse "gormTest/handlers/universe"
	serviceImage "gormTest/service/image"
	serviceTag "gormTest/service/tag"
	serviceUniverse "gormTest/service/universe"
)

// var validate = validator.New()
var imageHandler = handlerImage.NewImageHandler(serviceImage.NewImageService(dbImage.NewImageDatabase(database.ConnectToPostgreSQL())))
var tagHandler = handlerTag.NewTagHandler(serviceTag.NewTagService(dbTag.NewTagDatabase(database.ConnectToPostgreSQL())))
var universeHandler = handlerUniverse.NewUniverseHandler(serviceUniverse.NewUniverseService(dbUniverse.NewUniverseDatabase(database.ConnectToPostgreSQL())))

func UserRoute(e *echo.Echo) {

	//protected := e.Group("")
	//protected.Use(echojwt.WithConfig(config.NewConfig()))

	imageGroup := e.Group("/api/images")
	imageGroup.POST("", imageHandler.ImagePOST)
	imageGroup.GET("", imageHandler.ImagesGET)
	imageGroup.GET("/:id", imageHandler.ImageGetByID)
	imageGroup.PUT("/:id", imageHandler.ImagePUT)

	tagGroup := e.Group("/api/tags")
	tagGroup.POST("", tagHandler.TagPOST)
	tagGroup.GET("", tagHandler.TagGET)
	tagGroup.GET("/:id", tagHandler.TagGetByID)
	tagGroup.PUT("/:id", tagHandler.TagPUT)

	universeGroup := e.Group("/api/universe")
	universeGroup.POST("", universeHandler.UniversePOST)
	universeGroup.GET("", universeHandler.UniverseGET)
	universeGroup.GET("/:id", universeHandler.UniverseGetByID)
	universeGroup.PUT("/:id", universeHandler.UniversePUT)

	//protected.PUT("/user/:id", userHandler.EditUser)
	//protected.DELETE("/user/:id", userHandler.DeleteUser)
	//protected.POST("/user/:id/vote_for", userHandler.PostVoteFor)
	//protected.POST("/user/:id/vote_against", userHandler.PostVoteAgainst)
	//protected.PUT("/user/:id/vote", userHandler.ChangeVote)
	//protected.DELETE("/user/:id/vote", userHandler.DeleteVote)
	//
	//e.POST("/user", userHandler.CreateUser)
	//e.GET("/user/:id", userHandler.GetUser)
	//e.GET("/users", userHandler.GetAllUsers)
	//e.POST("/login", userHandler.Login)

}
