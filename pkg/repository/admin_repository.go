package repository

import (
	"os/user"

	"go.mongodb.org/mongo-driver/mongo"
)

// AdminRepo implements models.UserRepository
type AdminRepo struct {
	db *mongo.Client
}

// NewAdminRepo ..
func NewAdminRepo(db *mongo.Client) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}

// FindByID ..
func (r *AdminRepo) FindByID(ID int) (*user.User, error) {
	return &user.User{}, nil
}

// Save ..
func (r *AdminRepo) Save(user *user.User) error {
	return nil
}
