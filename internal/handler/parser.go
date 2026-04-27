package handler

import (
	"encoding/json"
	"net/http"

	"github.com/user/product-parser/internal/usecase"
)

type ParserHandler struct {
	useCase *usecase.ParseProductsUseCase
}

func NewParserHandler(useCase *usecase.ParseProductsUseCase) *ParserHandler {
	return &ParserHandler{useCase: useCase}
}

func (h *ParserHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	categorySlug := r.URL.Query().Get("category")
	if categorySlug == "" {
		http.Error(w, "category is required", http.StatusBadRequest)
		return
	}

	result, err := h.useCase.Execute(categorySlug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ParserHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.useCase.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}