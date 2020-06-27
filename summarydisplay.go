package main

import "log"

type SummaryDisplay struct {
	Temperatures   []float64
	Humidities     []float64
	Pressures      []float64
	AvgTemperature float64
	AvgHumidity    float64
	AvgPressure    float64
	WeatherData    Subject
}

// Constructor
func NewSummaryDisplay(weatherData Subject) *SummaryDisplay {
	sd := new(SummaryDisplay)
	sd.WeatherData = weatherData
	sd.WeatherData.RegisterObserver(sd)
	log.Println("Summary display initialized")
	return sd
}

func (sd *SummaryDisplay) update(
	temperature float64, humidity float64, pressure float64,
) {
	sd.Temperatures = append(sd.Temperatures, temperature)
	sd.Humidities = append(sd.Humidities, humidity)
	sd.Pressures = append(sd.Pressures, pressure)
	sd.AvgTemperature = average(sd.Temperatures)
	sd.AvgHumidity = average(sd.Humidities)
	sd.AvgPressure = average(sd.Pressures)
	sd.display()
}

func (sd *SummaryDisplay) display() {
	log.Printf(
		"Average - Temp %0.1fC, Humidity %0.1f%%, Pressure %0.1fhPa",
		sd.AvgTemperature, sd.AvgHumidity, sd.AvgPressure,
	)
}

func (sd *SummaryDisplay) watch() {
	log.Println("Current summary display is watching the weather")
}
