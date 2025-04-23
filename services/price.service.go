package services

import (
	"encoding/json"
	"epifigo/models"
	"fmt"
	"net/http"
	"os"
)

const baseURL = "https://finnhub.io/api/v1"
const symbol = "IBN"

var apiKey = os.Getenv("FINNHUB_API_KEY")

func GetCurrentPrice() (*models.Script, error) {
	url := fmt.Sprintf("%s/quote?symbol=%s&token=%s", baseURL, symbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quote models.Script
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		return nil, err
	}

	return &quote, nil
}

func DetectPattern() (string, string) {
	// Stub for now — Replace with real pattern logic
	return "Double Top (Simulated)", "Indicates a potential trend reversal from bullish to bearish."
}
