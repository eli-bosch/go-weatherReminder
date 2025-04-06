package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s,%s&appid=%s", City, State, Country, apiKey)
	fmt.Printf(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cordiateData CordinateResponse
	err = json.Unmarshal(body, &cordiateData)
	if err != nil {
		return nil, err
	}

	return &cordiateData, nil
}
