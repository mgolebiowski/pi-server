package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/mgolebiowski/pi-server/internal/models"
)

var (
	weatherCache    *models.ShortWeather
	lastWeatherCall time.Time
	weatherMutex    sync.Mutex
)

func GetWeather() (*models.ShortWeather, error) {
	weatherMutex.Lock()
	defer weatherMutex.Unlock()

	if weatherCache != nil && time.Since(lastWeatherCall) < 3600*time.Second {
		return weatherCache, nil
	}
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	u := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=50.0766907&lon=20.0125615&appid=%s&units=metric", apiKey)
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
