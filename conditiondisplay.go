package main

import "log"

type Observer interface {
	update(
		temperature float64, humidity float64, pressure float64,
	)
}

type DisplayElements interface {
	display()
}

type CurrentConditionsDisplay struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	WeatherData Subject
}

// Constructor
func NewCurrentConditionDisplay(weatherData Subject) *CurrentConditionsDisplay {
	ccd := new(CurrentConditionsDisplay)
	ccd.WeatherData = weatherData
	ccd.WeatherData.RegisterObserver(ccd)
	log.Println("Condition display initialized")
	return ccd
}

func (ccd *CurrentConditionsDisplay) update(
	temperature float64, humidity float64, pressure float64,
) {
	ccd.Temperature = temperature
	ccd.Humidity = humidity
	ccd.Pressure = pressure
	ccd.display()
}

func (ccd *CurrentConditionsDisplay) display() {
	log.Printf(
		"Current condition - Temp %0.1fC, Humidity %0.1f%%, Pressure %0.1fhPa",
		ccd.Temperature, ccd.Humidity, ccd.Pressure,
	)
}

func (ccd *CurrentConditionsDisplay) watch() {
	log.Println("Current condition display is watching the weather")
}
