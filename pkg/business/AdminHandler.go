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

	//w.Write([]byte(fmt.Sprintf("admin %v \n", admin)))
	utils.RespondJSON(w, 201, true, "", admin)
}

// AdminLogin , in this fucntion we will login as admin
func (handler *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {

	admin := database.Admin{}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		w.Write([]byte(fmt.Sprintf("enter username and password")))
		return
	}

	admin, _ = handler.adminRepo.AdminLogin(username, password)

	//encodedToken := utils.TokenExtractor(jwtauth.TokenFromHeader(r))
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": admin})
	w.Write([]byte(fmt.Sprintf("data %v \n", tokenString)))
}
