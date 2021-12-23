package controllers

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/services"
	"clean-architecture/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	user, err := u.service.GetOneUser(lib.ParseUUID(paramID))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.SetPaginationScope(utils.Paginate(c)).GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}

	limit := c.MustGet(constants.Limit).(int64)
	size := c.MustGet(constants.Page).(int64)

	c.JSON(
		200,
		gin.H{
			"data":       users["data"],
			"pagination": gin.H{"has_next": (users["count"].(int64) - limit*size) > 0, "count": users["count"]},
		})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	metadata, _ := c.MustGet(constants.File).(lib.UploadedFiles)
	files := metadata.GetMultipleFiles("files[]")
	for i := range files {
		fmt.Println(files[i])
	}

	if err := u.service.Create(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	paramID := lib.ParseUUID(c.Param("id"))

	user, err := u.service.GetOneUser(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.CustomBind(c.Request, &user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	metadata, _ := c.MustGet(constants.File).(lib.UploadedFiles)
	user.ProfilePic = lib.SignedURL(metadata.GetFile("file").URL)

	if err := u.service.UpdateUser(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	if err := u.service.DeleteUser(lib.ParseUUID(paramID)); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
