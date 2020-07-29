package business

import (
	"fmt"
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
	if user, err := h.categoryRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, Category"))
}
