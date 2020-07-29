package business

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
)

// UserHandler will hold everything that controller needs
type UserHandler struct {
	userRepo domain.UserRepository
}

// NewUserHandler returns a new BaseHandler
func NewUserHandler(userRepo domain.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// HelloUser returns Hello, World
func (h *UserHandler) HelloUser(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
	w.Write([]byte("Hello, User"))
}
