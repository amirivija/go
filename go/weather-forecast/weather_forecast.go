// Package weather package povides the weather condition at the specified city.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation is the location for which you need to fetch the current condition.
var CurrentLocation string

// Forecast returns the weather condition forecast for the input city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
