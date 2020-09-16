package repository

import (
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
