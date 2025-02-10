package permission

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// PermissionRepository database structure
type Repository struct {
	infrastructure.Database
	logger framework.Logger
}

// NewPermissionRepository creates a new Permission repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

func (r *Repository) GetAllPermission() ([]models.Permission, error) {
	var permissions []models.Permission
	query := r.DB.Model(&models.Permission{}).Find(&permissions)
	return permissions, query.Error
}

func (r *Repository) GetPermissionByID(id uint64) (models.Permission, error) {
	var permission models.Permission
	query := r.DB.Model(&models.Permission{}).Where("id=?", id).Find(&permission)
	return permission, query.Error
}
