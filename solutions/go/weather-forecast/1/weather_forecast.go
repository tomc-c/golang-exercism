// Package weather helps with weater forecast.
package weather

var (
    // CurrentCondition defines current conditions.
	CurrentCondition string
	// CurrentLocation defines current conditions.
    CurrentLocation  string
)
// Forecast returns weather forecast for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
