package business

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business/utils"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain/database"
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

// AdminRegister , in this fucntion we will login as admin
func (handler *AdminHandler) AdminRegister(w http.ResponseWriter, r *http.Request) {

	admin := database.Admin{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		w.Write([]byte(fmt.Sprintf("error")))
		return
	}
	defer r.Body.Close()

	handler.adminRepo.AdminRegister(&admin)
	utils.RespondJSON(w, 201, true, "", admin)
}

// AdminLogin , in this fucntion we will login as admin
func (handler *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {

	admin := database.Admin{}

	phone := r.FormValue("phone")
	password := r.FormValue("password")

	if phone == "" || password == "" {
		utils.RespondJSON(w, 400, false, "phone or password are empty", nil)
		return
	}

	admin, _ = handler.adminRepo.AdminLogin(phone, password)

	if admin.Phone == "" || admin.Password == "" {
		utils.RespondJSON(w, 403, false, "not found", nil)
		return
	}

	//encodedToken := utils.TokenExtractor(jwtauth.TokenFromHeader(r))
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"id": admin.ID})

	utils.RespondJSON(w, 200, true, "success", tokenString)

}
