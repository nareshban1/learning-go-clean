package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/types"
)

// UserService service layer
type Service struct {
	logger     framework.Logger
	repository Repository
}

// NewUserService creates a new userservice
func NewService(
	logger framework.Logger,
	userRepository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: userRepository,
	}
}

// Create creates the user in database
func (s Service) Create(user *models.User) error {
	return s.repository.Create(user).Error
}

// GetOneUser gets one user
func (s Service) GetUserByID(userID types.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "uuid = ?", userID).Error
}

// GetOneUser gets one user
func (s Service) DeleteUserByID(userID types.BinaryUUID) (err error) {
	return s.repository.DeleteUser(userID)
}

// GetRawUserFromID gets the raw user from id
func (r *Repository) GetRawUserFromID(userID uint) (user *models.User, err error) {
	r.logger.Info("[UserRepository...GetRawUserFromID]")

	query := r.Model(&models.User{}).Where("id = ?", userID).First(&user)

	return user, query.Error
}

// Get ALL USERS
func (s Service) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

// Update User
func (s Service) UpdateUser(userId types.BinaryUUID, user *models.User) (models.User, error) {
	return s.repository.UpdateUser(userId, user)
}
