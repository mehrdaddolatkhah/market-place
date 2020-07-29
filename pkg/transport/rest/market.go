package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// MarketRestAPI : handle all market routes
func MarketRestAPI(router *chi.Mux, marketHandler *business.MarketHandler) *chi.Mux {

	router.Get("/api/v1/market", marketHandler.HelloMarket)

	return router
}
