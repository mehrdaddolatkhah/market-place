package business

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
)

// AdminHandler will hold everything that controller needs
type AdminHandler struct {
	adminRepo domain.AdminRepository
}

// NewAdminHandler returns a new BaseHandler
func NewAdminHandler(adminRepo domain.AdminRepository) *AdminHandler {
	return &AdminHandler{
		adminRepo: adminRepo,
	}
}

// HelloAdmin returns Hello, World
func (h *AdminHandler) HelloAdmin(w http.ResponseWriter, r *http.Request) {
	if user, err := h.adminRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))

	w.Write([]byte("Hello, Admin"))
}
