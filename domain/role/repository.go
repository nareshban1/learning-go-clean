package role

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// UserRepository database structure
type Repository struct {
	infrastructure.Database
	logger framework.Logger
}

// NewUserRepository creates a new user repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

func (r *Repository) GetAllRole() ([]models.Role, error) {
	var roles []models.Role
	query := r.DB.Model(&models.Role{}).Preload("Permissions").Find(&roles)
	return roles, query.Error
}
