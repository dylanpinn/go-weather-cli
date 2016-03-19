package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body, err := getWeatherResponseBody()

	if err != nil {
		panic(err)
	}

	openWeather := OpenWeather{}
	err = json.Unmarshal(body, &openWeather)
	if err != nil {
		panic(err)
	}

	for i := range openWeather.List {
		fmt.Printf("\nWeather in %s is %.2f°C",
			openWeather.List[i].Name,
			openWeather.List[i].Weather.NormalisedCurrentTemp())
	}
}

func (w Weather) NormalisedCurrentTemp() float64 {
	return w.CurrentTemp - 273.15
}

func getWeatherResponseBody() ([]byte, error) {
	url := "http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.81&lon=144.96&cnt=10"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting weather: %v", err)
		return []byte(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading weather: %v", err)
		return []byte(""), err
	}

	return body, nil
}

type OpenWeather struct {
	List []City `json:"list"`
}

type City struct {
	Weather Weather `json:"main"`
	Name    string  `json:"name"`
}

type Weather struct {
	CurrentTemp float64 `json:"temp"`
	MaxTemp     float64 `json:"temp_max"`
}
