package repository

import (
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthRepo implements models.UserRepository
type AuthRepo struct {
	db *mongo.Client
}

// NewAuthRepo ..
func NewAuthRepo(db *mongo.Client) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

// Login ..
func (r *AuthRepo) Login(phone string) (string, error) {
	return "", nil
}

// Verify ..
func (r *AuthRepo) Verify(phone string, code string) (*database.User, error) {
	return nil, nil
}
