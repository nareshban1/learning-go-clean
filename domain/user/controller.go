package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"strconv"

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

	userID, err := strconv.ParseUint(paramID, 10, 64)
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

type PermissionsResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type UserResponse struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	RoleID      uint                  `json:"roleId"`
	RoleName    string                `json:"roleName"`
	Permissions []PermissionsResponse `json:"permissions"`
}

// gelAllUsers
func (u *Controller) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}
	var usersResponse []UserResponse
	for _, value := range users {
		var permissions []PermissionsResponse

		for _, permission := range value.Role.Permissions {
			permissions = append(permissions, PermissionsResponse{
				ID:   permission.ID,
				Name: permission.Name,
			})
		}

		usersResponse = append(usersResponse, UserResponse{
			ID:          value.ID,
			Name:        value.FirstName + ` ` + value.LastName,
			Email:       value.Email,
			RoleID:      value.RoleID,
			RoleName:    value.Role.Name,
			Permissions: permissions,
		})
	}

	c.JSON(200, gin.H{
		"data": usersResponse,
	})
}

// GetOneUser gets one user
func (u *Controller) GetUserByID(c *gin.Context) {
	paramID := c.Param("id")
	var userResponse UserResponse
	userID, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	user, err := u.service.GetUserByID(userID)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	var permissions []PermissionsResponse

	for _, permission := range user.Role.Permissions {
		permissions = append(permissions, PermissionsResponse{
			ID:   permission.ID,
			Name: permission.Name,
		})
	}

	userResponse = UserResponse{
		ID:          user.ID,
		Name:        user.FirstName + ` ` + user.LastName,
		RoleID:      user.RoleID,
		RoleName:    user.Role.Name,
		Permissions: permissions,
	}

	c.JSON(200, gin.H{
		"data": userResponse,
	})

}

// GetOneUser gets one user
func (u *Controller) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := strconv.ParseUint(paramID, 10, 64)
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
