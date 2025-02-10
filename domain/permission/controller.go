package permission

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

// permissionController data type
type Controller struct {
	service *Service
	logger  framework.Logger
	env     *framework.Env
}

// NewPermissionController creates new permission controller
func NewController(
	permissionService *Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		service: permissionService,
		logger:  logger,
		env:     env,
	}
}

func (u *Controller) CreatePermission(c *gin.Context) {
	var permission models.Permission

	if err := c.Bind(&permission); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	// check if the permission already exists

	if err := u.service.Create(&permission); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "Permission created"})
}

func (u *Controller) GetAllPermission(c *gin.Context) {
	permissions, err := u.service.GetAllPermissions()
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": permissions})
}

func (u *Controller) GetPermissionByID(c *gin.Context) {
	permissionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}
	permissions, err := u.service.GetPermissionByID(permissionID)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": permissions})
}
