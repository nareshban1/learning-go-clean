package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/types"
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

// ExistsByEmail checks if the user exists by email
func (r *Repository) ExistsByEmail(email string) (bool, error) {
	r.logger.Info("[UserRepository...Exists]")

	users := make([]models.User, 0, 1)
	query := r.DB.Where("email = ?", email).Limit(1).Find(&users)

	return query.RowsAffected > 0, query.Error
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	r.logger.Info("[UserRepository...Exists]")

	var users []models.User
	query := r.DB.Model(&models.User{}).Preload("Role.Permissions").Find(&users)
	return users, query.Error
}

func (r *Repository) UpdateUser(userID types.BinaryUUID, user *models.User) (models.User, error) {
	r.logger.Info("[UserRepository...Exists]")
	var userData models.User
	query := r.DB.Model(&models.User{}).Where("uuid = ?", userID).Updates(user).Find(&userData)
	return userData, query.Error
}

func (r *Repository) DeleteUser(userID types.BinaryUUID) error {
	r.logger.Info("[UserRepository...Exists]")

	query := r.DB.Model(&models.User{}).Where("uuid = ?", userID).Delete(&models.User{})
	return query.Error
}
