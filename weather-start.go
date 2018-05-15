package main

import ("fmt"
"encoding/json"
)

type DataPoint struct {
Time uint64 `json:"time"`
Summary string `json:"summary"`
Icon string `json:"icon"`
PrecipIntensity float64 `json:"precipIntensity"`
PrecipProbability float64 `json:"precipProbability"`
Temperature float64 `json:"temperature"`
ApparentTemperature float64 `json:"apparentTemperature"`
Dewpoint float64 `json:"dewPoint"`
Humidity float64 `json:"humidity"`
Pressure float64 `json:"pressure"`
WindSpeed float64 `json:"windSpeed"`
WindGust float64 `json:"windGust"`
WindBearing int `json:"windBearing"`
CloudCover float64 `json:"cloudCover"`
UvIndex int `json:"uvIndex"`
Visibility float64 `json:"visibility"`
Ozone float64 `json:"ozone"`
}

const Data string = `{"time":1526414400,
						"summary":"Partly Cloudy",
						"icon":"partly-cloudy-night",
						"precipIntensity":0,
						"precipProbability":0.76,
						"temperature":58.78,
						"apparentTemperature":58.78,
						"dewPoint":50.64,
						"humidity":0.74,
						"pressure":1021.92,
						"windSpeed":10.66,
						"windGust":20.13,
						"windBearing":312,
						"cloudCover":0.51,
						"uvIndex":0,
						"visibility":9.55,
						"ozone":336.45}`

func main() {
	var wd DataPoint
	err := json.Unmarshal([]byte(Data), &wd)
	fmt.Println(err)
	fmt.Println(wd.PrecipProbability)
}

