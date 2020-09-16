package business

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business/utils"
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
	encodedToken := utils.TokenExtractor(jwtauth.TokenFromHeader(r))
	w.Write([]byte(fmt.Sprintf("data %v \n", encodedToken)))
}
