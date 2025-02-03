package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mgolebiowski/pi-server/internal/models"
	"github.com/mgolebiowski/pi-server/internal/ttss"
	"github.com/mgolebiowski/pi-server/internal/weather"
)

func main() {
	ttss.InitTripsCache()

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		trams, err := ttss.GetStop()
		if err != nil {
			http.Error(w, "Failed to fetch tram data", http.StatusInternalServerError)
		}

		weather, err := weather.GetWeather()
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		}

		response, err := json.Marshal(models.StopResponse{Trams: trams, Weather: *weather})
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(response)
	})

	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
