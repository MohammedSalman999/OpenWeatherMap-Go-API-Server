package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin  float64 `json:"temp"`
		Celsius float64 // Add this field to hold the Celsius temperature
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go \n"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, fmt.Errorf("error loading API config: %w", err)
	}

	// Validate city name
	if city == "" {
		return weatherData{}, errors.New("empty city name provided")
	}

	// Build API request URL
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s",
		apiConfig.OpenWeatherMapApiKey, city)

	// Send request
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, fmt.Errorf("error sending API request: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return weatherData{}, fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Parse JSON response
	var data weatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return weatherData{}, fmt.Errorf("error parsing JSON response: %w", err)
	}

	// Convert temperature to Celsius
	data.Main.Celsius = data.Main.Kelvin - 273.15

	return data, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json, charset = utf-8")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8000", nil)
}