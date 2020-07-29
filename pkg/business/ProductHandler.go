package business

import (
	"fmt"
	"net/http"

	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
)

// ProductHandler will hold everything that controller needs
type ProductHandler struct {
	productRepo domain.ProductRepository
}

// NewProductHandler returns a new BaseHandler
func NewProductHandler(productRepo domain.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
	}
}

// HelloProduct returns Hello, World
func (h *ProductHandler) HelloProduct(w http.ResponseWriter, r *http.Request) {
	if user, err := h.productRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, Product"))
}
