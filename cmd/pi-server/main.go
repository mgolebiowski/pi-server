package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mgolebiowski/pi-server/models"
)

var (
	weatherCache    *models.ShortWeather
	lastWeatherCall time.Time
	weatherMutex    sync.Mutex
)

func getWeather() (*models.ShortWeather, error) {
	weatherMutex.Lock()
	defer weatherMutex.Unlock()

	if weatherCache != nil && time.Since(lastWeatherCall) < 3600*time.Second {
		return weatherCache, nil
	}

	u := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=50.0766907&lon=20.0125615&appid=%s&units=metric", os.Getenv("OPEN_WEATHER_API_KEY"))
	weatherResp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer weatherResp.Body.Close()

	weatherBody, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		return nil, err
	}

	var weather models.WeatherResponse
	err = json.Unmarshal(weatherBody, &weather)
	if err != nil {
		return nil, err
	}

	weatherCache = &models.ShortWeather{
		Temperature: int(weather.Main.Temp),
		Icon:        weather.Weather[0].Icon,
	}
	lastWeatherCall = time.Now()

	return weatherCache, nil
}

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://www.ttss.krakow.pl/internetservice/services/passageInfo/stopPassages/stop?stop=407&mode=departure")
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		var passages models.StopPassages
		err = json.Unmarshal(body, &passages)
		if err != nil {
			http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		}

		var trams []models.Tram
		for _, passage := range passages.Actual {
			if passage.ActualRelativeTime > 5*60 {
				newEta := strings.Replace(passage.MixedTime, "%UNIT_MIN%", "minut", 1)
				trams = append(trams, models.Tram{
					Line:      passage.PatternText,
					Direction: passage.Direction,
					ETA:       newEta,
				})
			}
		}

		weather, err := getWeather()

		tramsJSON, err := json.Marshal(models.StopResponse{Trams: trams, Weather: *weather})

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(tramsJSON)
	})

	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
