package models

type Passage struct {
	ActualRelativeTime int    `json:"actualRelativeTime"`
	ActualTime         string `json:"actualTime"`
	Direction          string `json:"direction"`
	MixedTime          string `json:"mixedTime"`
	PassageID          string `json:"passageid"`
	PatternText        string `json:"patternText"`
	PlannedTime        string `json:"plannedTime"`
	RouteID            string `json:"routeId"`
	Status             string `json:"status"`
	TripID             string `json:"tripId"`
	VehicleID          string `json:"vehicleId"`
}

type Route struct {
	Alerts     []interface{} `json:"alerts"`
	Authority  string        `json:"authority"`
	Directions []string      `json:"directions"`
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	RouteType  string        `json:"routeType"`
	ShortName  string        `json:"shortName"`
}

type StopPassages struct {
	Actual           []Passage     `json:"actual"`
	Directions       []interface{} `json:"directions"`
	FirstPassageTime int64         `json:"firstPassageTime"`
	GeneralAlerts    []interface{} `json:"generalAlerts"`
	LastPassageTime  int64         `json:"lastPassageTime"`
	Old              []Passage     `json:"old"`
	Routes           []Route       `json:"routes"`
	StopName         string        `json:"stopName"`
	StopShortName    string        `json:"stopShortName"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type WeatherResponse struct {
	Weather []Weather `json:"weather"`
	Base    string    `json:"base"`
	Main    Main      `json:"main"`
}

type StopResponse struct {
	Trams   []Tram       `json:"trams"`
	Weather ShortWeather `json:"weather"`
}

type ShortWeather struct {
	Temperature int    `json:"temperature"`
	Icon        string `json:"icon"`
}

type Tram struct {
	Line      string `json:"line"`
	Direction string `json:"direction"`
	ETA       string `json:"eta"`
	ToCenter  bool   `json:"isTripToCityCenter"`
}
