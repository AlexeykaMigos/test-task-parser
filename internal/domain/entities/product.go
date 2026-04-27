package entities

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

type Store struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ParseResult struct {
	Category   string    `json:"category"`
	ParsedAt   string    `json:"parsed_at"`
	TotalCount int       `json:"total_count"`
	Products   []Product `json:"products"`
}