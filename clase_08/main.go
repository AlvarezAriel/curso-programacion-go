package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
	ID                string    `json:"id"`
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

func main() {
	recolectarDataMeli()
}

func recolectarDataMeli() {
	searchResult := makeRequest("https://api.mercadolibre.com/sites/MLA/search?q=RTX%203080")

	csvfile, err := os.Create("test.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)

	for _, resultItem := range searchResult.Results {
		_ = csvwriter.Write([]string{
			resultItem.ID,
			resultItem.Title,
			strconv.Itoa(resultItem.Price),
			strconv.Itoa(resultItem.SoldQuantity),
		})
	}

	csvwriter.Flush()

	fmt.Println(searchResult)
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
