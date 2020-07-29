package repository

import (
	"os/user"

	"go.mongodb.org/mongo-driver/mongo"
)

// MarketRepo implements models.UserRepository
type MarketRepo struct {
	db *mongo.Client
}

// NewMarketRepo ..
func NewMarketRepo(db *mongo.Client) *MarketRepo {
	return &MarketRepo{
		db: db,
	}
}

// FindByID ..
func (r *MarketRepo) FindByID(ID int) (*user.User, error) {
	return &user.User{}, nil
}

// Save ..
func (r *MarketRepo) Save(user *user.User) error {
	return nil
}
