package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// PublicRouter : handle all user routes
func PublicRouter(router chi.Router, userHandler *business.AuthHandler) chi.Router {

	// Login for admin
	router.Post("/admin/login", userHandler.Login)

	// Login and Verify for marketer
	router.Post("/market/login", userHandler.Verify)
	router.Post("/market/verify", userHandler.Verify)

	// Login and Verify for normal users
	router.Post("/login", userHandler.Login)
	router.Post("/verify", userHandler.Verify)

	return router
}
