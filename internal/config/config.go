package config

import "os"

type Config struct {
	BaseURL   string
	ProxyURL  string
	StoreID   string
	UserAgent string
}

func Load() *Config {
	return &Config{
		BaseURL:   getEnv("BASE_URL", "https://kuper.ru"),
		ProxyURL:  getEnv("PROXY_URL", ""),
		StoreID:   getEnv("STORE_ID", ""),
		UserAgent: getEnv("USER_AGENT", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}[+] up 3/3
 ✔ Image random-parser      Built                                                                                                           3.4s ✔ Network random_default   Created                                                                                                         0.1s ✔ Container product-parser Created                                                                                                         0.1sAttaching to product-parser
product-parser  | === Product Parser ===
product-parser  | 
product-parser  |   1. Молочные продукты
product-parser  |   2. Хлеб и выпечка
product-parser  |   3. Овощи и фрукты
product-parser  |   4. Мясо и птица
product-parser  |   5. Рыба и морепродукты
product-parser  |   6. Сыры
product-parser  |   7. Колбасы
product-parser  |   8. Крупы и макароны
product-parser  |   9. Чай и кофе
product-parser  |   10. Напитки
product-parser  | 
product-parser  | 
product-parser  | Parsing: Молочные продукты
product-parser  | 2026/04/27 16:57:18 Error: failed to get products: parse error: invalid character '<' looking for beginning of value
product-parser  | 
product-parser  | Parsing: Хлеб и выпечка
product-parser  | 
product-parser  | === Done ===
product-parser  | 2026/04/27 16:57:18 Error: failed to get products: parse error: invalid character '<' looking for beginning of value
product-parser exited with code 0

	return defaultValue
}