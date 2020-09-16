package business

import (
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
