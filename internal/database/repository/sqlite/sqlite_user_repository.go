package sqlite

import (
	"github.com/emejotaw/users-api-bff/internal/database/entity"
	"gorm.io/gorm"
)

type SqliteUserRepository struct {
	db *gorm.DB
}

func NewSqliteUserRepository(db *gorm.DB) *SqliteUserRepository {

	return &SqliteUserRepository{
		db: db,
	}
}

func (r *SqliteUserRepository) Create(user *entity.User) error {

	return r.db.Create(user).Error
}

func (r *SqliteUserRepository) FindByID(userID string) (*entity.User, error) {

	user := &entity.User{}
	err := r.db.First(&user, userID).Error
	return user, err
}

func (r *SqliteUserRepository) FindAll() ([]entity.User, error) {

	users := []entity.User{}
	err := r.db.Find(&users).Error
	return users, err
}
