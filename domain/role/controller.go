package role

import (
	"clean-architecture/domain/models"
	"clean-architecture/domain/permission"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service           *Service
	permissionService *permission.Service
	logger            framework.Logger
	env               *framework.Env
}

// NewUserController creates new user controller
func NewController(
	roleService *Service,
	permissionService *permission.Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		permissionService: permissionService,
		service:           roleService,
		logger:            logger,
		env:               env,
	}
}

type roleCreateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Permissions []uint64 `json:"permissions" binding:"required"`
}

func (u *Controller) CreateRole(c *gin.Context) {
	var roleRequest roleCreateRequest

	if err := c.Bind(&roleRequest); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	var permissions []models.Permission
	for _, id := range roleRequest.Permissions {
		perm, err := u.permissionService.GetPermissionByID(id)
		if err != nil {
			responses.HandleError(u.logger, c, err)
			return
		}
		permissions = append(permissions, perm)
	}

	role := models.Role{
		Name:        roleRequest.Name,
		Permissions: permissions,
	}
	u.logger.Info(role, "ROLE DATAAAA")

	if err := u.service.Create(&role); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "Role created"})
}

func (u *Controller) GetRoles(c *gin.Context) {
	roles, err := u.service.GetAllRoles()
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": roles})
}
