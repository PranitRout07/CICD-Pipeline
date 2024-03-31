package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiURL string

func TestQuerry(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/data/2.5/weather" {
			t.Errorf("Expected path /data/2.5/weather, got %s", r.URL.Path)
		}

		mockResponse := `{
			"name": "London",
			"main": {
				"temp": 280
			},
			"weather": [
				{"main": "Clouds", "description": "scattered clouds"}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer ts.Close()

	oldURL := apiURL
	defer func() { apiURL = oldURL }()
	apiURL = ts.URL

	city := "London"
	data, err := querry(city)
	if err != nil {
		t.Errorf("Error querying weather data: %v", err)
	}

	if data.Name != city {
		t.Errorf("Expected city name %s, got %s", city, data.Name)
	}
	if len(data.Weather) == 0 {
		t.Error("Expected weather data, got none")
	}
	// if int(data.Main.Kelvin) != 7 {
	// 	t.Errorf("Expected temperature 7°C, got %f°C", data.Main.Kelvin)
	// }
}




