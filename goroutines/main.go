package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const APIKey = "37307e910be70c5cab09dcf85137a2a3"

type Data struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) Data {
	var data Data

	defer wg.Done()
	fmt.Println(city, " Start!")
	defer fmt.Println(city, " End!")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, APIKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)

	}

	ch <- fmt.Sprintf("Weather in %s is -> %.2fc", city, data.Main.Temp)

	return data
}

func main() {
	start := time.Now()

	cities := []string{"Ermesinde", "London", "Copenhagen", "Tokyo", "Oslo", "Berlin", "Paris", "New York", "San Francisco"}

	ch := make(chan string)

	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(start))
}
