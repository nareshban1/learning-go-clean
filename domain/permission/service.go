package permission

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
)

// permissionService service layer
type Service struct {
	logger     framework.Logger
	repository Repository
}

// NewPermissionService creates a new
func NewService(
	logger framework.Logger,
	permissionRepository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: permissionRepository,
	}
}

// Create creates the permission in database
func (s Service) Create(permission *models.Permission) error {
	return s.repository.Create(permission).Error
}

func (s Service) GetAllPermissions() ([]models.Permission, error) {
	return s.repository.GetAllPermission()
}

func (s Service) GetPermissionByID(id uint64) (models.Permission, error) {
	return s.repository.GetPermissionByID(id)
}
