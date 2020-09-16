package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// UserRouter : handle all user routes
func UserRouter(router chi.Router, userHandler *business.UserHandler) chi.Router {

	//router.Get("/api/v1/user", userHandler.HelloUser)

	return router
}
