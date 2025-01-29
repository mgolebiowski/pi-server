package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mgolebiowski/pi-server/internal/models"
	"github.com/mgolebiowski/pi-server/internal/ttss"
	"github.com/mgolebiowski/pi-server/internal/weather"
)

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://ttss.krakow.pl/internetservice/services/passageInfo/stopPassages/stop?stop=407&mode=departure")
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
				isToCityCenter, err := ttss.IsTripToCityCenter(passage.TripID)
				if err != nil {
					http.Error(w, "Failed to check if trip is to city center", http.StatusInternalServerError)
				}
				trams = append(trams, models.Tram{
					Line:      passage.PatternText,
					Direction: passage.Direction,
					ETA:       newEta,
					ToCenter:  isToCityCenter,
				})
			}
		}

		weather, err := weather.GetWeather()
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		}

		tramsJSON, err := json.Marshal(models.StopResponse{Trams: trams, Weather: *weather})
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(tramsJSON)
	})

	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
