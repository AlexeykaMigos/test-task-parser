package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Product struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Unit   string  `json:"unit,omitempty"`
	URL    string  `json:"url"`
	Brand  string  `json:"brand,omitempty"`
}

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
}

var categories = []Category{
	{ID: "1", Name: "Молочные продукты", Slug: "moloko-syr-yayca/molochnye-programmy"},
	{ID: "2", Name: "Хлеб и выпечка", Slug: "khleb-i-vypechka"},
	{ID: "3", Name: "Овощи и фрукты", Slug: "ovoshchi-frukty-zelen"},
	{ID: "4", Name: "Мясо и птица", Slug: "myaso-ptitsa"},
	{ID: "5", Name: "Рыба и морепродукты", Slug: "ryba-i-moreprodukty"},
	{ID: "6", Name: "Сыры", Slug: "moloko-syr-yayca/syry"},
	{ID: "7", Name: "Колбасы", Slug: "myaso-ptitsya/kolbasy"},
	{ID: "8", Name: "Крупы и макароны", Slug: "krupy-kakarony"},
	{ID: "9", Name: "Чай и кофе", Slug: "chay-kofe-kakao"},
	{ID: "10", Name: "Напитки", Slug: "napitki"},
}

func main() {
	config := Config{
		StoreID:   getEnv("STORE_ID", ""),
		ProxyURL:  getEnv("PROXY_URL", ""),
		BaseURL:   "https://kuper.ru",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	}

	args := os.Args[1:]
	if len(args) > 0 && args[0] != "" {
		config.ProxyURL = args[0]
	}

	fmt.Println("=== Product Parser for Kuper.ru ===")
	fmt.Println()

	fmt.Printf("Using %d categories\n\n", len(categories))
	for i, cat := range categories {
		fmt.Printf("  %d. %s\n", i+1, cat.Name)
	}
	fmt.Println()

	catNums := getEnv("CATEGORIES", "1,2")
	categoryIDs := strings.Split(catNums, ",")

	outputDir := getEnv("OUTPUT_DIR", ".")

	for _, catNum := range categoryIDs {
		var idx int
		fmt.Sscanf(catNum, "%d", &idx)
		idx--
		if idx < 0 || idx >= len(categories) {
			fmt.Printf("Category %d not found, skipping...\n", idx+1)
			continue
		}

		cat := categories[idx]
		fmt.Printf("\n=== Parsing: %s ===\n", cat.Name)

		products, err := getProducts(config, cat.Slug)
		if err != nil {
			log.Printf("Error getting products for %s: %v", cat.Name, err)
			continue
		}

		filename := fmt.Sprintf("products_%s.json", strings.ReplaceAll(cat.Slug, "/", "_"))
		saveProducts(products, filename, cat.Name, outputDir)

		fmt.Printf("Saved %d products to %s\n", len(products), filename)
		for i, p := range products {
			if i < 5 {
				fmt.Printf("  - %s | %.2f₽ | %s\n", p.Name, p.Price, p.URL)
			}
		}
		if len(products) > 5 {
			fmt.Printf("  ... and %d more\n", len(products)-5)
		}
	}

	fmt.Println("\n=== Parsing Complete ===")
}

func getProducts(config Config, categorySlug string) ([]Product, error) {
	var allProducts []Product
	page := 1
	limit := 24

	for {
		apiURL := fmt.Sprintf("%s/api/catalog/category/%s?page=%d&limit=%d",
			config.BaseURL, categorySlug, page, limit)

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", config.UserAgent)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en;q=0.8")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Referer", config.BaseURL+"/")

		client := createClient(config)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		body, err := readResponseBody(resp)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		var prodResp struct {
			Products []ProductItem `json:"products"`
			Meta     struct {
				TotalPages int `json:"totalPages"`
			} `json:"meta"`
		}

		if err := json.Unmarshal(body, &prodResp); err != nil {
			var altResp struct {
				Data struct {
					Products []ProductItem `json:"products"`
					Meta     struct {
						TotalPages int `json:"totalPages"`
					} `json:"meta"`
				} `json:"data"`
			}
			if err := json.Unmarshal(body, &altResp); err != nil {
				return nil, fmt.Errorf("parse error: %w", err)
			}
			prodResp.Products = altResp.Data.Products
			prodResp.Meta.TotalPages = altResp.Data.Meta.TotalPages
		}

		if len(prodResp.Products) == 0 {
			break
		}

		for _, p := range prodResp.Products {
			product := Product{
				Name:  p.Name,
				Price: p.Price,
				Unit:  p.Unit,
				URL:   config.BaseURL + "/product/" + p.Slug,
			}
			allProducts = append(allProducts, product)
		}

		if page >= prodResp.Meta.TotalPages {
			break
		}
		page++

		time.Sleep(300 * time.Millisecond)
	}

	return allProducts, nil
}

type ProductItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
	Brand string  `json:"brand"`
}

type Config struct {
	StoreID   string
	ProxyURL  string
	BaseURL   string
	UserAgent string
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	encoding := resp.Header.Get("Content-Encoding")

	switch encoding {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	default:
		return io.ReadAll(resp.Body)
	}
}

func createClient(config Config) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}

	if config.ProxyURL != "" {
		proxyURL, err := url.Parse(config.ProxyURL)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func saveProducts(products []Product, filename string, categoryName string, outputDir string) {
	type ExportData struct {
		Category   string    `json:"category"`
		ParsedAt   time.Time `json:"parsed_at"`
		TotalCount int       `json:"total_count"`
		Products   []Product `json:"products"`
	}

	export := ExportData{
		Category:   categoryName,
		ParsedAt:   time.Now(),
		TotalCount: len(products),
		Products:   products,
	}

	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	filepath := outputDir + "/" + filename
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		log.Printf("Error writing file %s: %v", filepath, err)
	}
}
