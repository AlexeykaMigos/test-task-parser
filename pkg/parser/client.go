package parser

import (
	"fmt"
	"time"
)

type Client struct {
	baseURL   string
	proxyURL  string
	storeID   string
	userAgent string
	demo      bool
}

func NewClient(baseURL, proxyURL, storeID, userAgent string) *Client {
	return &Client{
		baseURL:   baseURL,
		proxyURL:  proxyURL,
		storeID:   storeID,
		userAgent: userAgent,
		demo:      true,
	}
}

func (c *Client) GetProducts(categorySlug string) ([]Product, error) {
	time.Sleep(500 * time.Millisecond)

	return getDemoProducts(categorySlug), nil
}

func (c *Client) GetCategories() ([]Category, error) {
	return defaultCategories(), nil
}

func (c *Client) GetStores() ([]Store, error) {
	return defaultStores(), nil
}

func getDemoProducts(categorySlug string) []Product {
	products := map[string][]Product{
		"moloko-syr-yayca/molochnye-programmy": {
			{Name: "Молоко Простоквашино 3.2%", Price: 89.99, Unit: "л", Slug: "moloko-prostokvashino-32", Brand: "Простоквашино"},
			{Name: "Кефир Домик в деревне 1%", Price: 65.50, Unit: "л", Slug: "kefir-domik-v-derevne-1", Brand: "Домик в деревне"},
			{Name: "Творог Савушкин 5%", Price: 129.00, Unit: "г", Slug: "tvorog-savushkin-5", Brand: "Савушкин"},
			{Name: "Йогурт Активиа 1%", Price: 75.99, Unit: "шт", Slug: "yogurt-aktiviya-1", Brand: "Активиа"},
			{Name: "Сметана Простоквашино 15%", Price: 95.00, Unit: "г", Slug: "smetana-prostokvashino-15", Brand: "Простоквашино"},
			{Name: "Ряженка Домик в деревне 4%", Price: 72.50, Unit: "л", Slug: "ryajenka-domik-4", Brand: "Домик в деревне"},
			{Name: "Молоко Parmalat 3.5%", Price: 110.00, Unit: "л", Slug: "moloko-parmalat-35", Brand: "Parmalat"},
			{Name: "Творожок Агуша 4.5%", Price: 55.00, Unit: "шт", Slug: "tvorojok-agusha-45", Brand: "Агуша"},
		},
		"khleb-i-vypechka": {
			{Name: "Хлеб Бородинский", Price: 45.00, Unit: "шт", Slug: "hleb-borodinsky", Brand: "Бородинский"},
			{Name: "Хлеб белый пшеничный", Price: 35.00, Unit: "шт", Slug: "hleb-bely-pshenichny", Brand: ""},
			{Name: "Батон нарезной", Price: 42.00, Unit: "шт", Slug: "baton-nareznoi", Brand: ""},
			{Name: "Лаваш тонкий", Price: 55.00, Unit: "шт", Slug: "lavash-tonky", Brand: ""},
			{Name: "Пирог с капустой", Price: 180.00, Unit: "шт", Slug: "pirog-s-kapustoy", Brand: ""},
			{Name: "Сосиски в тесте", Price: 85.00, Unit: "шт", Slug: "sosiski-v-teste", Brand: ""},
			{Name: "Булочка сдобная", Price: 28.00, Unit: "шт", Slug: "bulochka-sdobnaya", Brand: ""},
			{Name: "Круассан с маслом", Price: 65.00, Unit: "шт", Slug: "kruassan-s-maslom", Brand: ""},
		},
		"ovoshchi-frukty-zelen": {
			{Name: "Яблоки Гренни Смит", Price: 149.00, Unit: "кг", Slug: "yabloki-grenni-smith", Brand: ""},
			{Name: "Бананы Эквадор", Price: 95.00, Unit: "кг", Slug: "banany-ekvador", Brand: ""},
			{Name: "Апельсины Испания", Price: 120.00, Unit: "кг", Slug: "apelsiny-ispaniya", Brand: ""},
			{Name: "Помидоры красные", Price: 180.00, Unit: "кг", Slug: "pomidory-krasnye", Brand: ""},
			{Name: "Огурцы тепличные", Price: 150.00, Unit: "кг", Slug: "ogurtsy-teplichnye", Brand: ""},
			{Name: "Картофель молодой", Price: 65.00, Unit: "кг", Slug: "kartofel-molodoy", Brand: ""},
			{Name: "Морковь Россия", Price: 55.00, Unit: "кг", Slug: "morkov-rossiya", Brand: ""},
			{Name: "Капуста белокочанная", Price: 45.00, Unit: "кг", Slug: "kapusta-belokochannaya", Brand: ""},
		},
		"myaso-ptitsa": {
			{Name: "Курица тушка", Price: 220.00, Unit: "кг", Slug: "kuritsa-tushka", Brand: ""},
			{Name: "Свинина вырезка", Price: 350.00, Unit: "кг", Slug: "svinina-vyrezka", Brand: ""},
			{Name: "Говядина стейк", Price: 550.00, Unit: "кг", Slug: "govyadina-steyk", Brand: ""},
			{Name: "Куриное филе", Price: 280.00, Unit: "кг", Slug: "kurinoe-file", Brand: ""},
			{Name: "Бедро куриное", Price: 250.00, Unit: "кг", Slug: "bedro-kurinoe", Brand: ""},
			{Name: "Фарш свиной", Price: 300.00, Unit: "кг", Slug: "farsh-svinoy", Brand: ""},
			{Name: "Шашлык свиной", Price: 420.00, Unit: "кг", Slug: "shashlyk-svinoy", Brand: ""},
			{Name: "Индейка филе", Price: 450.00, Unit: "кг", Slug: "indeyka-file", Brand: ""},
		},
		"ryba-i-moreprodukty": {
			{Name: "Сельдь солёная", Price: 180.00, Unit: "кг", Slug: "seld-solenaia", Brand: ""},
			{Name: "Сёмга слабосолёная", Price: 850.00, Unit: "кг", Slug: "semga-slabosolenaya", Brand: ""},
			{Name: "Минтай свежий", Price: 180.00, Unit: "кг", Slug: "mintay-svejiy", Brand: ""},
			{Name: "Креветки варёные", Price: 450.00, Unit: "кг", Slug: "krevetki-varenye", Brand: ""},
			{Name: "Филе трески", Price: 380.00, Unit: "кг", Slug: "file-treski", Brand: ""},
			{Name: "Горбуша свежая", Price: 320.00, Unit: "кг", Slug: "gorbusha-svejaia", Brand: ""},
		},
		"syry": {
			{Name: "Сыр Российский", Price: 450.00, Unit: "кг", Slug: "syr-rossiysky", Brand: ""},
			{Name: "Сыр Адыгейский", Price: 380.00, Unit: "кг", Slug: "syr-adygeysky", Brand: ""},
			{Name: "Сыр Пармезан", Price: 1200.00, Unit: "кг", Slug: "syr-parmezhan", Brand: ""},
			{Name: "Сыр Брынза", Price: 350.00, Unit: "кг", Slug: "syr-brynza", Brand: ""},
			{Name: "Сыр Фета", Price: 550.00, Unit: "кг", Slug: "syr-feta", Brand: ""},
		},
	}

	if prods, ok := products[categorySlug]; ok {
		result := make([]Product, len(prods))
		for i, p := range prods {
			result[i] = Product{
				ID:    fmt.Sprintf("%d", i+1),
				Name:  p.Name,
				Slug:  p.Slug,
				Price: p.Price,
				Unit:  p.Unit,
				Brand: p.Brand,
				URL:   "https://kuper.ru/product/" + p.Slug,
			}
		}
		return result
	}

	return []Product{
		{Name: "Товар демо 1", Price: 99.99, Unit: "шт", Slug: "demo-1", URL: "https://kuper.ru/product/demo-1"},
		{Name: "Товар демо 2", Price: 149.50, Unit: "кг", Slug: "demo-2", URL: "https://kuper.ru/product/demo-2"},
		{Name: "Товар демо 3", Price: 75.00, Unit: "л", Slug: "demo-3", URL: "https://kuper.ru/product/demo-3"},
	}
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
	Brand string  `json:"brand"`
	URL   string  `json:"url"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Store struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func defaultCategories() []Category {
	return []Category{
		{ID: "1", Name: "Молочные продукты", Slug: "moloko-syr-yayca/molochnye-programmy"},
		{ID: "2", Name: "Хлеб и выпечка", Slug: "khleb-i-vypechka"},
		{ID: "3", Name: "Овощи и фрукты", Slug: "ovoshchi-frukty-zelen"},
		{ID: "4", Name: "Мясо и птица", Slug: "myaso-ptitsa"},
		{ID: "5", Name: "Рыба и морепродукты", Slug: "ryba-i-moreprodukty"},
		{ID: "6", Name: "Сыры", Slug: "syry"},
		{ID: "7", Name: "Колбасы", Slug: "myaso-ptitsya/kolbasy"},
		{ID: "8", Name: "Крупы и макароны", Slug: "krupy-kakarony"},
		{ID: "9", Name: "Чай и кофе", Slug: "chay-kofe-kakao"},
		{ID: "10", Name: "Напитки", Slug: "napitki"},
	}
}

func defaultStores() []Store {
	return []Store{
		{ID: "1052", Name: "Москва, Тверская", Address: "г. Москва, ул. Тверская, д. 1"},
		{ID: "1053", Name: "Москва, Ленинградская", Address: "г. Москва, Ленинградское ш., д. 5"},
		{ID: "1054", Name: "СПб, Невский", Address: "г. Санкт-Петербург, Невский пр., д. 10"},
	}
}