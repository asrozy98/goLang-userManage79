package router

import (
	"goLang-userManage79/handler"
	"goLang-userManage79/repository"
	"goLang-userManage79/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Word",
		})
	})
	userRepository := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	user := route.Group("/user")
	user.POST("/", userHandler.CreateUser)
	user.GET("/", userHandler.GetUsers)
	user.GET("/:id", userHandler.GetUser)
	user.POST("/:id", userHandler.UpdateUser)
	user.DELETE("/:id", userHandler.DeleteUser)

	route.Run()
}
