package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// AdminRouter todo : change it later. useless
// AdminRouter : handle all admin routes
func AdminRouter(router chi.Router, adminHandler *business.AdminHandler) http.Handler {
	router.Get("/hello", adminHandler.AdminLogin)
	router.Post("/register", adminHandler.AdminRegister)
	return router
}
