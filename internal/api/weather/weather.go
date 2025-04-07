package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
	//Figure this out from openweather api
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
}

func FetchWeather(Longitude float64, Latitude float64) (*WeatherResponse, error) {
	apiKey := os.Getenv("WEATHER_API_KEY") //Make name more descriptive
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.2f&lon=%.2f&units=imperial&appid=%s", Latitude, Longitude, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
