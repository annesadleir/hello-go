package main

type Weather struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Hours     Hourly  `json:"hourly"`
	Offset    int     `json:"offset"`
}

type Hourly struct {
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	Data    []DataPoint `json:"data"`
}

type DataPoint struct {
	Time                uint64  `json:"time"`
	Summary             string  `json:"summary"`
	Icon                string  `json:"icon"`
	PrecipIntensity     float64 `json:"precipIntensity"`
	PrecipProbability   float64 `json:"precipProbability"`
	Temperature         float64 `json:"temperature"`
	ApparentTemperature float64 `json:"apparentTemperature"`
	Dewpoint            float64 `json:"dewPoint"`
	Humidity            float64 `json:"humidity"`
	Pressure            float64 `json:"pressure"`
	WindSpeed           float64 `json:"windSpeed"`
	WindGust            float64 `json:"windGust"`
	WindBearing         int     `json:"windBearing"`
	CloudCover          float64 `json:"cloudCover"`
	UvIndex             int     `json:"uvIndex"`
	Visibility          float64 `json:"visibility"`
	Ozone               float64 `json:"ozone"`
}
