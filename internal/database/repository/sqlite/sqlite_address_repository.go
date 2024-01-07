package sqlite

import (
	"github.com/emejotaw/users-api-bff/internal/database/entity"
	"gorm.io/gorm"
)

type SqliteAddressRepository struct {
	db *gorm.DB
}

func NewSqliteAddressRepository(db *gorm.DB) *SqliteAddressRepository {
	return &SqliteAddressRepository{
		db: db,
	}
}

func (r *SqliteAddressRepository) FindByUserID(userID string) (*entity.Address, error) {

	address := &entity.Address{}

	err := r.db.Where("user_id = ?", userID).Find(address).Error
	return address, err
}
