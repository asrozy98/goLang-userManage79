package handler

import (
	"fmt"
	"goLang-userManage79/model"
	"goLang-userManage79/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (handler *userHandler) CreateUser(c *gin.Context) {
	var userRequest model.UsersRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(400, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := handler.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (handler *userHandler) GetUsers(c *gin.Context) {
	users, err := handler.userService.GetUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

func (handler *userHandler) UpdateUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var userRequest model.UsersRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(400, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := handler.userService.UpdateUser(id, userRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (handler *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := handler.userService.GetUser(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (handler *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := handler.userService.DeleteUser(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully deleted user"})
}
