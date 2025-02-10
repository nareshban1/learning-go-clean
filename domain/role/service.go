package role

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"

	"gorm.io/gorm"
)

// UserService service layer
type Service struct {
	logger     framework.Logger
	repository Repository
}

// NewUserService creates a new
func NewService(
	logger framework.Logger,
	roleRepository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: roleRepository,
	}
}

// Create creates the user in database
func (s Service) Create(role *models.Role) error {

	return s.repository.Session(&gorm.Session{FullSaveAssociations: true}).Create(role).Error
}

func (s Service) GetAllRoles() ([]models.Role, error) {
	return s.repository.GetAllRole()
}
