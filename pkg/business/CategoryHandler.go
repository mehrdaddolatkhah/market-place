package business

import (
	"net/http"

	"github.com/mehrdaddolatkhah/cafekala_server/pkg/domain"
)

// CategoryHandler will hold everything that controller needs
type CategoryHandler struct {
	categoryRepo domain.CategoryRepository
}

// NewCategoryHandler returns a new BaseHandler
func NewCategoryHandler(categoryRepo domain.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: categoryRepo,
	}
}

// HelloCategory returns Hello, World
func (h *CategoryHandler) HelloCategory(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, Category"))
}
