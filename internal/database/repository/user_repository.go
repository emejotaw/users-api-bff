package repository

import "github.com/emejotaw/users-api-bff/internal/database/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(userID string) (*entity.User, error)
	FindAll() ([]entity.User, error)
}
