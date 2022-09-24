package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func createWeatherList(data [][]string) []WeatherRecord {
	// convert csv lines to array of structs
	var weatherList []WeatherRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec WeatherRecord
			for j, field := range line {
				if j == 0 {
					rec.Time = field
				} else if j == 1 {
					var err error
					rec.Tavg, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 2 {
					var err error
					rec.Tmin, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 3 {
					var err error
					rec.Tmax, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 4 {
					var err error
					rec.Prcp, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 5 {
					var err error
					rec.Snow, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 6 {
					var err error
					rec.Wdir, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 7 {
					var err error
					rec.Wpgt, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				} else if j == 8 {
					var err error
					rec.Pres, err = strconv.ParseFloat(field, 32)
					if err != nil {
						continue
					}
				}
			}
			weatherList = append(weatherList, rec)
		}
	}
	return weatherList
}

func readWeatherList(file string) []WeatherRecord {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return createWeatherList(data)
}
