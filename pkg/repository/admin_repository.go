package repository

import (
	"time"

	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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

// AdminRegister ...
func (r *AdminRepo) AdminRegister(admin *database.Admin) (*mongo.InsertOneResult, error) {

	adminCollection := r.db.Database(DatabaseName).Collection(database.ADMIN_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := adminCollection.InsertOne(ctx, admin)
	return result, nil
}

// AdminLogin ...
func (r *AdminRepo) AdminLogin(phone string, code string) (*database.Admin, error) {

	return nil, nil
}
