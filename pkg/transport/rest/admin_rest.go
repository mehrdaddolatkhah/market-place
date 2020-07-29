package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// AdminRestAPI : handle all admin routes
func AdminRestAPI(router *chi.Mux, adminHandler *business.AdminHandler) *chi.Mux {

	router.Get("/api/v1/admin", adminHandler.HelloAdmin)

	return router
}
