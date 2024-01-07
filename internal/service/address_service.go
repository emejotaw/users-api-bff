package service

import (
	"log"

	"github.com/emejotaw/users-api-bff/internal/database/entity"
	"github.com/emejotaw/users-api-bff/internal/database/repository"
	"github.com/emejotaw/users-api-bff/internal/database/repository/sqlite"
	"gorm.io/gorm"
)

type AddressService struct {
	repository repository.AddressRepository
}

func NewAddressService(db *gorm.DB) *AddressService {

	repository := sqlite.NewSqliteAddressRepository(db)

	return &AddressService{
		repository: repository,
	}
}

func (s *AddressService) FindByUserID(userID string) (*entity.Address, error) {

	log.Printf("Looking for address of user id %s", userID)

	return s.repository.FindByUserID(userID)
}
