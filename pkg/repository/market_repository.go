package repository

import (
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
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

// MarketerLogin ...
func (r *MarketRepo) MarketerLogin(phone string) (string, error) {

	return "", nil
}

// MarketerVerify ...
func (r *MarketRepo) MarketerVerify(phone string, code string) (*database.Market, error) {
	return nil, nil
}
