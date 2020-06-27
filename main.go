package main

func main() {
	weatherData := new(WeatherData)
	currentDisplay := NewCurrentConditionDisplay(weatherData)
	summaryDisplay := NewSummaryDisplay(weatherData)
	currentDisplay.watch()
	summaryDisplay.watch()

	weatherData.SetMeasurements(30, 30, 1000)
	weatherData.SetMeasurements(27, 40, 1002)
	weatherData.SetMeasurements(29, 30, 1000)
	weatherData.SetMeasurements(26, 40, 960)
	weatherData.SetMeasurements(30, 36, 1020)
	weatherData.SetMeasurements(29, 60, 1002)
}
