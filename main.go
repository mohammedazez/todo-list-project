package main

import (
	"todoAPIGolang/config"
	"todoAPIGolang/handler"
	"todoAPIGolang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	r := gin.Default()

	r.GET("/users", userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", userHandler.UpdateUserByIDHandler)

	// r.POST("/users/login", userHandler.CreateUserHandler)

	r.Run(":4444")
}
