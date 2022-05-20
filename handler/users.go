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
		errorMessages := []any{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(400, gin.H{
			"success": false,
			"message": "Validation error",
			"error":   errorMessages,
		})
		return
	}

	user, err := handler.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User created",
		"data":    user,
	})
}

func (handler *userHandler) GetUsers(c *gin.Context) {
	pageString := c.Query("page")
	page, _ := strconv.Atoi(pageString)
	if page == 0 {
		page = 1
	}

	limitString := c.Query("limit")
	limit, _ := strconv.Atoi(limitString)
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit
	users, err, total := handler.userService.GetUsers(offset, limit)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success":      true,
		"data":         users,
		"perPageCount": len(users),
		"allCount":     total,
	})
}

func (handler *userHandler) UpdateUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var userRequest model.UsersRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorMessages := []any{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   errorMessages,
		})
		return
	}

	user, err := handler.userService.UpdateUser(id, userRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User updated",
		"data":    user,
	})
}

func (handler *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := handler.userService.GetUser(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func (handler *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := handler.userService.DeleteUser(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User deleted",
	})
}
