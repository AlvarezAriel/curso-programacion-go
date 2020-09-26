package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SearchResult struct {
	SiteID     string       `json:"site_id"`
	Query      string       `json:"query"`
	Pagination Paging       `json:"paging"`
	Results    []ItemResult `json:"results"`
}

type Paging struct {
	Total          int `json:"total"`
	PrimaryResults int `json:"primary_results"`
	Offset         int `json:"offset"`
	Limit          int `json:"limit"`
}

type ItemResult struct {
	MeliId            string    `json:"id"`
	Title             string    `json:"title"`
	Price             int       `json:"price"`
	CurrencyID        string    `json:"currency_id"`
	AvailableQuantity int       `json:"available_quantity"`
	SoldQuantity      int       `json:"sold_quantity"`
	BuyingMode        string    `json:"buying_mode"`
	ListingTypeID     string    `json:"listing_type_id"`
	StopTime          time.Time `json:"stop_time"`
	Condition         string    `json:"condition"`
}

type ItemModel struct {
	gorm.Model
	ItemResult
	Puesto int
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	recolectarDataMeli()
}

func mainEjemplo() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	newProduct := Product{Code: "D42", Price: 100}
	db.Create(&newProduct)

	// Read
	var product Product
	db.First(&product, newProduct.ID)     // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	db.Select(&product, "code = ?", "D42")

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	db.Delete(&product, 1)
}

func recolectarDataMeli() {
	db, err := gorm.Open(sqlite.Open("meli.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ItemModel{})

	//db.Exec("UPDATE item_models SET puesto=id WHERE puesto IS NULL")

	pageOffset := 0
	for i := 0; i < 4; i++ {
		requestUrl := "https://api.mercadolibre.com/sites/MLA/search?q=zapatillas&offset=" + strconv.Itoa(pageOffset)
		fmt.Println(requestUrl)

		searchResult := makeRequest(requestUrl)
		fmt.Println("Guardando " + strconv.Itoa(len(searchResult.Results)) + " resultados.")
		for index, resultItem := range searchResult.Results {
			item := ItemModel{ItemResult: resultItem}
			item.Puesto = 1 + index + pageOffset
			db.Create(&item)
		}

		pageOffset += 50

		time.Sleep(time.Millisecond * 500)
	}
}

func makeRequest(searchRequestUrl string) SearchResult {
	resp, err := http.Get(searchRequestUrl)

	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	searchResult := SearchResult{}

	json.Unmarshal(body, &searchResult)
	return searchResult
}
