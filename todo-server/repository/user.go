package repository

import (
	"beards.ly/todo/model"
	"gorm.io/gorm"
)

// UserRepository repo
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository for wire
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

// FindAll users
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Debug().Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

// FindByID user
func (r *UserRepository) FindByID(id int) (model.User, error) {
	var user model.User
	err := r.DB.Debug().Model(&model.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// FindByUsername user
func (r *UserRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Debug().Model(&model.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Save user
func (r *UserRepository) Save(user model.User) (model.User, error) {
	err := r.DB.Debug().Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
