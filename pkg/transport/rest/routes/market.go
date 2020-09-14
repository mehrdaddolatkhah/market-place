package rest

import (
	"github.com/go-chi/chi"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
)

// MarketRouter : handle all market routes
func MarketRouter(router chi.Router, marketHandler *business.MarketHandler) chi.Router {

	router.Get("/api/v1/market", marketHandler.HelloMarket)

	return router
}
