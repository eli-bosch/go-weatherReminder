package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type CordinateResponse struct {
	//Figure this out from openweather api
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

func FetchState(City string, State string, Country string) (*CordinateResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&limit=1&appid=%s", City, State, Country, apiKey)
	fmt.Println()

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cordiateData []CordinateResponse
	err = json.Unmarshal(body, &cordiateData)
	if err != nil {
		return nil, err
	}

	if len(cordiateData) == 0 {
		return nil, fmt.Errorf("no location data found for %s, %s, %s", City, State, Country)
	}

	return &cordiateData[0], nil
}
