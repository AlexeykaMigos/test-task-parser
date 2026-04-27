package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/user/product-parser/internal/config"
	"github.com/user/product-parser/internal/repository"
	"github.com/user/product-parser/internal/usecase"
	"github.com/user/product-parser/pkg/parser"
)

func main() {
	cfg := config.Load()

	categories := flag.String("categories", "1,2", "Category numbers (comma-separated)")
	proxy := flag.String("proxy", "", "Proxy URL")
	output := flag.String("output", ".", "Output directory")
	flag.Parse()

	if *proxy != "" {
		cfg.ProxyURL = *proxy
	}

	client := parser.NewClient(cfg.BaseURL, cfg.ProxyURL, cfg.StoreID, cfg.UserAgent)

	productRepo := repository.NewProductRepo(client)
	categoryRepo := repository.NewCategoryRepo(client)

	useCase := usecase.NewParseProductsUseCase(productRepo, categoryRepo)

	fmt.Println("=== Product Parser ===\n")

	cats, err := useCase.GetCategories()
	if err != nil {
		log.Printf("Error getting categories: %v", err)
	}

	for i, cat := range cats {
		fmt.Printf("  %d. %s\n", i+1, cat.Name)
	}
	fmt.Println()

	categoryIDs := strings.Split(*categories, ",")

	for _, catNum := range categoryIDs {
		var idx int
		fmt.Sscanf(catNum, "%d", &idx)
		idx--
		if idx < 0 || idx >= len(cats) {
			continue
		}

		cat := cats[idx]
		fmt.Printf("\nParsing: %s\n", cat.Name)

		result, err := useCase.Execute(cat.Slug)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		filename := fmt.Sprintf("products_%s.json", strings.ReplaceAll(cat.Slug, "/", "_"))
		saveResult(result, *output+"/"+filename)

		fmt.Printf("Saved %d products to %s\n", result.TotalCount, filename)
		for i, p := range result.Products {
			if i < 3 {
				fmt.Printf("  - %s | %.2f\n", p.Name, p.Price)
			}
		}
	}

	fmt.Println("\n=== Done ===")
}

func saveResult(result interface{}, filepath string) {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Printf("Error marshaling: %v", err)
		return
	}
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		log.Printf("Error writing file: %v", err)
	}
}