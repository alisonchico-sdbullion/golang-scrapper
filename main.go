package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func scrapeQuotes() ([]Quote, error) {
    var quotes []Quote

    c := colly.NewCollector(
        colly.AllowedDomains("quotes.toscrape.com"),
        colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
    )

    c.OnHTML(".quote", func(e *colly.HTMLElement) {
        quote := Quote{
            Text:   e.ChildText(".text"),
            Author: e.ChildText(".author"),
        }
        quotes = append(quotes, quote)
    })

    for i := 1; i <= 10 && len(quotes) < 100; i++ {
        url := "http://quotes.toscrape.com/page/" + strconv.Itoa(i) + "/"
        log.Printf("Visiting: %s", url)
        if err := c.Visit(url); err != nil {
            return nil, err
        }
    }

    if len(quotes) > 100 {
        quotes = quotes[:100]
    }

    return quotes, nil
}


func quotesHandler(w http.ResponseWriter, r *http.Request) {
	quotes, err := scrapeQuotes()
	if err != nil {
		http.Error(w, "Failed to fetch quotes", http.StatusInternalServerError)
		log.Printf("Error scraping quotes: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		http.Error(w, "Failed to encode quotes", http.StatusInternalServerError)
		log.Printf("Error encoding quotes: %v", err)
	}
}

func main() {
	http.HandleFunc("/quotes", quotesHandler)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}