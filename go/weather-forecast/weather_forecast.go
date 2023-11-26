// Package weather provids weather conditions for Goblinocus cities.
package weather

// CurrentCondition represents weather at a point in time.
var CurrentCondition string
// CurrentLocation represents location of the city whose weather we are currently working on.
var CurrentLocation string

// Forecast displays current weather for given Goblinocus city in human readable and repeatable format.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
