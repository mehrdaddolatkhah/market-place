package repository

import (
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
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

// AdminLogin ...
func (r *AdminRepo) AdminLogin(phone string, code string) (*database.Admin, error) {

	return nil, nil
}
