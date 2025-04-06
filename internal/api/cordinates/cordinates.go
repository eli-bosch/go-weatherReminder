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
}

type openWeatherCordResponse struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func FetchState(City string, State string, Country string) (*CordinateResponse, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s,%s&appid=%s", City, State, Country, apiKey)
	fmt.Println("Requesting URL:", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var cordinateData openWeatherCordResponse
	err = json.Unmarshal(body, &cordinateData)
	if err != nil {
		return nil, err
	}

	var responseData CordinateResponse

	responseData.Latitude = cordinateData.Coord.Lat
	responseData.Longitude = cordinateData.Coord.Lon
	responseData.Country = cordinateData.Sys.Country
	responseData.Name = cordinateData.Name

	return &responseData, nil
}
