package main

import "log"

// observer_pattern will implement observer pattern
// from head first design pattern

// Subject is an interface
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

// WeatherData
type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	observers   []Observer
}

func (w *WeatherData) MeasurementChanged() {
	log.Println("Measurement has changed!")
	w.NotifyObserver()
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(
				w.observers[:i],
				w.observers[i+1:]...,
			)
		}
	}
}
func (w *WeatherData) NotifyObserver() {
	for i, observer := range w.observers {
		observer.update(
			w.Temperature,
			w.Humidity,
			w.Pressure,
		)
		log.Printf("%dth observer updated", i+1)
	}
}

func (w *WeatherData) SetMeasurements(
	temperature float64, humidity float64, pressure float64,
) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure
	w.MeasurementChanged()
}
