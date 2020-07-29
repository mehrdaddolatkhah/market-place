package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// UserRestAPI : handle all user routes
func UserRestAPI(router *chi.Mux, userHandler *business.UserHandler) *chi.Mux {

	router.Get("/api/v1/user", userHandler.HelloUser)

	return router
}
