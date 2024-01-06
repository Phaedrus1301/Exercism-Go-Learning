// Package weather provides tools that can help with
// forecasting the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents certain type of weather.
var CurrentCondition string

// CurrentLocation represents Location within the country of Goblinocus.
var CurrentLocation string

// Forecast returns a string value which helps display what type of weather is currently present in the mentioned location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
