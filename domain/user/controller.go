package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/types"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service *Service
	logger  framework.Logger
	env     *framework.Env
}

// NewUserController creates new user controller
func NewController(
	userService *Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		service: userService,
		logger:  logger,
		env:     env,
	}
}

// CreateUser creates the new user
func (u *Controller) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	// check if the user already exists

	if err := u.service.Create(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// Update User
func (u *Controller) UpdateUser(c *gin.Context) {
	var user models.User
	paramID := c.Param("id")
	u.logger.Error(paramID)

	userID, err := types.ShouldParseUUID(paramID)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	if _, err := u.service.GetUserByID(userID); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	if err := c.Bind(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	// check if the user already exists
	userUpdateData, err := u.service.UpdateUser(userID, &user)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": userUpdateData})
}

// gelAllUsers
func (u *Controller) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}

// GetOneUser gets one user
func (u *Controller) GetUserByID(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	user, err := u.service.GetUserByID(userID)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetOneUser gets one user
func (u *Controller) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	if _, err := u.service.GetUserByID(userID); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	if err := u.service.DeleteUserByID(userID); err != nil {
		u.logger.Error(err)
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": "User Deleted",
	})

}
