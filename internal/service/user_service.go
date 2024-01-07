package service

import (
	"github.com/emejotaw/users-api-bff/graph/model"
	"github.com/emejotaw/users-api-bff/internal/database/entity"
	"github.com/emejotaw/users-api-bff/internal/database/repository"
	"github.com/emejotaw/users-api-bff/internal/database/repository/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {

	repostory := sqlite.NewSqliteUserRepository(db)
	return &UserService{
		repository: repostory,
	}
}

func (s *UserService) Create(userInput model.UserInput) (*entity.User, error) {

	user := &entity.User{
		ID:         uuid.New().String(),
		Name:       userInput.Name,
		Email:      userInput.Email,
		DocumentId: userInput.DocumentID,
		Address: &entity.Address{
			ID:           uuid.New().String(),
			State:        userInput.Address.State,
			City:         userInput.Address.City,
			Neighborhood: userInput.Address.Neighborhood,
			Street:       userInput.Address.Street,
			Number:       userInput.Address.Number,
			ZipCode:      userInput.Address.ZipCode,
		},
	}

	err := s.repository.Create(user)
	return user, err
}

func (s *UserService) FindAll() ([]entity.User, error) {
	return s.repository.FindAll()
}
