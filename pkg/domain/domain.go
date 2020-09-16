package domain

import (
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
}

type ProductRepository interface {
}

type MarketRepository interface {
	MarketerLogin(phone string) (string, error)
	MarketerVerify(phone string, code string) (*database.Market, error)
}

type AdminRepository interface {
	AdminLogin(phone string, password string) (*database.Admin, error)
	AdminRegister(*database.Admin) (*mongo.InsertOneResult, error)
}

type CategoryRepository interface {
}

type AuthRepository interface {
	Login(phone string) (string, error)
	Verify(phone string, code string) (*database.User, error)
}
