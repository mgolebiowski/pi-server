package ttss

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/mgolebiowski/pi-server/internal/models"
)

func GetStop() ([]models.Tram, error) {
	resp, err := http.Get("https://ttss.krakow.pl/internetservice/services/passageInfo/stopPassages/stop?stop=407&mode=departure")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var passages models.StopPassages
	err = json.Unmarshal(body, &passages)
	if err != nil {
		return nil, err
	}

	var trams []models.Tram
	for _, passage := range passages.Actual {
		// We need at least 5 minutes to get to the stop
		if passage.ActualRelativeTime > 5*60 {
			// "5 %UNIT_MIN%" ->"5 min"
			newEta := strings.Replace(passage.MixedTime, "%UNIT_MIN%", "min", 1)
			toCenter, err := IsTripToCityCenter(passage.TripID)
			if err != nil {
				return nil, err
			}
			trams = append(trams, models.Tram{
				Line:      passage.PatternText,
				Direction: passage.Direction,
				ETA:       newEta,
				ToCenter:  toCenter,
			})
		}
	}

	return trams, nil
}
