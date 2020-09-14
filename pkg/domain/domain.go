package domain

import (
	"os/user"
)

// type Repository interface {

// User
// RegisterUser(input database.User) (string, error)
// SendSmsVerification(userCredential Credentials) (string, error)
// GetSmsVerification(userData UserData) (*map[string]string, string)
// ResetPasswordSendMobile(userCredential Credentials) string
// ResetPasswordGetOtp(userData UserData) string
// UserLogin(userCredential Credentials, password string) (*map[string]string, string)
// GetAllUsers(userUUID uuid.UUID, user User) ([]User, string)
// FindUser(mobile string, email string, userUUID uuid.UUID) (*User, string)
// UpdateUser(user User, userUUID uuid.UUID) string
// DeleteUser(user User, userUUID uuid.UUID) string
// ChangeUserState(user User, adminUUID uuid.UUID) string

// // Product
// AddProduct(product Product, adminUUID uuid.UUID) (string, error)
// GetAllProducts(adminUUID uuid.UUID) ([]Product, string)
// UpdateProduct(product Product, adminId uuid.UUID) string
// ChangeProductState(product Product, adminId uuid.UUID) string
// GetProductsByName(name string, adminId uuid.UUID) ([]Product, string)
// GetProductsByCategory(categoryId string, adminId uuid.UUID) ([]Product, string)
// }

type UserRepository interface {
	FindByID(ID int) (*user.User, error)
	Save(user *user.User) error
}

type ProductRepository interface {
	FindByID(ID int) (*user.User, error)
	Save(user *user.User) error
}

type MarketRepository interface {
	FindByID(ID int) (*user.User, error)
	Save(user *user.User) error
}

type AdminRepository interface {
	FindByID(ID int) (*user.User, error)
	Save(user *user.User) error
}

type CategoryRepository interface {
	FindByID(ID int) (*user.User, error)
	Save(user *user.User) error
}

type AuthRepository interface {
	Login(phone string) (string, error)
	Verify(phone string, code string) (*user.User, error)
}
