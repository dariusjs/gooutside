package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCityTemperature(t *testing.T) {
	sampleResponse := `{"coord":{"lon":4.63,"lat":52.37},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"base":"stations","main":{"temp":18.55,"feels_like":16.4,"temp_min":17.78,"temp_max":19.44,"pressure":1024,"humidity":88},"visibility":10000,"wind":{"speed":6.2,"deg":210},"clouds":{"all":40},"dt":1599987897,"sys":{"type":1,"id":1524,"country":"NL","sunrise":1599973986,"sunset":1600020106},"timezone":7200,"id":2755002,"name":"Gemeente Bubblegum","cod":200}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, sampleResponse)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", responseData)
	cityWeather := CityWeather{}
	json.Unmarshal([]byte(responseData), &cityWeather)
	cityWeather.Name = strings.ReplaceAll((cityWeather.Name), " ", "_")

	openweatherApiKey := "12345"
	openweatherApi := ts.URL
	city := "Amsterdam"
	testCityWeather := getCityTemperature(openweatherApiKey, openweatherApi, city)
	fmt.Println(testCityWeather)

	// Our Expected Data
	expectedData := CityWeather{"Gemeente_Bubblegum", Data{18.55, 1024, 88}}

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode, "expected 200 status code")
	assert.Equal(t, testCityWeather, expectedData, "expected 200 status code")

}
