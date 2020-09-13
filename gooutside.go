package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type CityWeather struct {
	Name string `json:"name"`
	Data Data   `json:"main"`
}

type Data struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

func getConfig() (string, string) {
	openweatherApiKey, err := os.LookupEnv("OPENWEATHER_API_KEY")
	if err == false {
		log.Fatal("Environment variable OPENWEATHER_API_KEY is missing ", err)
	}
	influxDbAddress, err := os.LookupEnv("INFLUX_DB_ADDRESS")
	if err == false {
		log.Fatal("Environment variable OPENWEATHER_API_KEY is missing ", err)
	}
	return openweatherApiKey, influxDbAddress
}

func getCityTemperature(openweatherApiKey string, openweatherApi string, city string) CityWeather {
	openweatherUrl := openweatherApi + "/weather?q=" + city + "&units=metric&appid=" + openweatherApiKey
	response, err := http.Get(openweatherUrl)
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	cityWeather := CityWeather{}
	json.Unmarshal([]byte(responseData), &cityWeather)
	cityWeather.Name = strings.ReplaceAll((cityWeather.Name), " ", "_")
	return cityWeather
}

func postToInflux(cityWeather CityWeather, influxDbAddress string) {
	weatherData := "openweathermap," + "city=" + fmt.Sprint(cityWeather.Name) + " temperature=" + fmt.Sprint(cityWeather.Data.Temp) + ",pressure=" + fmt.Sprint(cityWeather.Data.Pressure) + ",humidity=" + fmt.Sprint(cityWeather.Data.Humidity)
	fmt.Println(weatherData)

	response, err := http.Post(influxDbAddress, "application/octet-stream", bytes.NewBuffer([]byte(weatherData)))
	if err != nil {
		log.Print(err)
	}
	fmt.Println(response)
}

func main() {
	openweatherApiKey, influxDbAddress := getConfig()
	openweatherApi := "http://api.openweathermap.org/data/2.5"
	city := "Haarlem"

	webserver := http.NewServeMux()
	cityWeather := getCityTemperature(openweatherApiKey, openweatherApi, city)
	postToInflux(cityWeather, influxDbAddress)
	tick := time.Tick(60 * time.Minute)
	for range tick {
		cityWeather := getCityTemperature(openweatherApiKey, openweatherApi, city)
		postToInflux(cityWeather, influxDbAddress)
	}
	err := http.ListenAndServe(":4001", webserver)
	log.Fatal(err)
}
