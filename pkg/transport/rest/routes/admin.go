package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// AdminRouter : handle all admin routes
func AdminRouter(router chi.Router, adminHandler *business.AdminHandler) http.Handler {
	router.Get("/hello", adminHandler.HelloAdmin)
	return router
}
