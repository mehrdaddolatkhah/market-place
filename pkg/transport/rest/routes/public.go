package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// PublicRouter : handle all user routes
func PublicRouter(router chi.Router, userHandler *business.AuthHandler, adminHandler *business.AdminHandler) chi.Router {

	// Login for admin
	router.Post("/login/admin", adminHandler.AdminLogin)

	// Login and Verify for marketer
	router.Post("/login/market", userHandler.Verify)
	router.Post("/verify/market", userHandler.Verify)

	// Login and Verify for normal users
	router.Post("/login", userHandler.Login)
	router.Post("/verify", userHandler.Verify)

	return router
}
