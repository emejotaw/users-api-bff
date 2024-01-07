package repository

import "github.com/emejotaw/users-api-bff/internal/database/entity"

type AddressRepository interface {
	FindByUserID(addressID string) (*entity.Address, error)
}
