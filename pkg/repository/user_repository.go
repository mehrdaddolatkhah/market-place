package repository

import (
	"os/user"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *mongo.Client
}

// NewUserRepo ..
func NewUserRepo(db *mongo.Client) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByID ..
func (r *UserRepo) FindByID(ID int) (*user.User, error) {
	return &user.User{}, nil
}

// Save ..
func (r *UserRepo) Save(user *user.User) error {
	return nil
}
