package business

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
)

// MarketHandler will hold everything that controller needs
type MarketHandler struct {
	marketRepo domain.MarketRepository
}

// NewMarketHandler returns a new BaseHandler
func NewMarketHandler(marketRepo domain.MarketRepository) *MarketHandler {
	return &MarketHandler{
		marketRepo: marketRepo,
	}
}

// HelloMarket returns Hello, World
func (h *MarketHandler) HelloMarket(w http.ResponseWriter, r *http.Request) {
	if user, err := h.marketRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	fmt.Println(jwtauth.FromContext(r.Context()))
	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))

	w.Write([]byte("Hello, Market"))
}
