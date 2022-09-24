package main

type WeatherRecord struct {
	Time string  `json:"time"`
	Tavg float64 `json:"tavg"`
	Tmin float64 `json:"tmin"`
	Tmax float64 `json:"tmax"`
	Prcp float64 `json:"prcp"`
	Snow float64 `json:"snow"`
	Wdir float64 `json:"wdir"`
	Wpgt float64 `json:"wpgt"`
	Pres float64 `json:"pres"`
}
