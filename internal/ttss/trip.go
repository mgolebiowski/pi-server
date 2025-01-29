package ttss

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stop struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type Passage struct {
	ActualTime string `json:"actualTime"`
	Status     string `json:"status"`
	Stop       Stop   `json:"stop"`
	StopSeqNum string `json:"stop_seq_num"`
}

type TripResponse struct {
	Actual        []Passage `json:"actual"`
	DirectionText string    `json:"directionText"`
	Old           []Passage `json:"old"`
	RouteName     string    `json:"routeName"`
}

func IsTripToCityCenter(tripID string) (bool, error) {
	u := fmt.Sprintf("https://ttss.krakow.pl/internetservice/services/tripInfo/tripPassages?tripId=%s&mode=departure", tripID)
	resp, err := http.Get(u)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var trip TripResponse
	err = json.NewDecoder(resp.Body).Decode(&trip)
	if err != nil {
		return false, err
	}

	for i, passage := range trip.Actual {
		if passage.Stop.ShortName == "407" {
			if trip.Actual[i+1].Stop.ShortName == "112" {
				return true, nil
			}
			return false, nil
		}
	}
	return false, nil
}
