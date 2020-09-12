package main

import (
	"net/http"
	"testing"
)

func MyRouter() http.Handler {
	r := http.NewServeMux()
	// r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	return r
}
func TestDoStuffWithTestServer(t *testing.T) {
	// server := httptest.NewServer(MyRouter())
	// defer server.Close()
	// resp, err := http.Get(server)

	// // Use Client & URL from our local test server
	// // api := API{server.Client(), server.URL}
	// openweatherApiKey := "12345"
	// openweatherApi := "http://test.api.weather"
	// city := "Amsterdam"

	// cityWeather := getCityTemperature(openweatherApiKey, openweatherApi, city)

	// // ok(t, err)
	// // equals(t, []byte("OK"), body)
	// assert.NoError(t, err)
	// assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")

}
