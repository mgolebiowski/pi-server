package ttss

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
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

const (
	CzyzynyStopID          = "407"
	StellaSawickiegoStopID = "112"
)

type tripCache struct {
	direction bool
	timestamp time.Time
}

var (
	cache = make(map[string]tripCache)
	mutex sync.RWMutex
	// Cache entries valid for 10 minutes
	cacheDuration = 10 * time.Minute
)

func IsTripToCityCenter(tripID string) (bool, error) {
	// Check cache first
	mutex.RLock()
	if cached, exists := cache[tripID]; exists {
		if time.Since(cached.timestamp) < cacheDuration {
			mutex.RUnlock()
			return cached.direction, nil
		}
	}
	mutex.RUnlock()

	// Make API call if not in cache or expired
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

	// Determine direction
	result := false
	for i, passage := range trip.Actual {
		if passage.Stop.ShortName == CzyzynyStopID {
			if i+1 < len(trip.Actual) && trip.Actual[i+1].Stop.ShortName == StellaSawickiegoStopID {
				result = true
			}
			break
		}
	}

	// Store in cache
	mutex.Lock()
	cache[tripID] = tripCache{
		direction: result,
		timestamp: time.Now(),
	}
	mutex.Unlock()

	return result, nil
}
