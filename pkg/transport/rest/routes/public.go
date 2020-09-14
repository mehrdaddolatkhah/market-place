package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// PublicRouter : handle all user routes
func PublicRouter(router chi.Router, userHandler *business.AuthHandler) chi.Router {

	router.Post("/login", userHandler.Login)

	router.Post("/verify", userHandler.Verify)

	return router
}
