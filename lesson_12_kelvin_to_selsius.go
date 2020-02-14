package main

import (
	"fmt"
)

func kelvinToCelsius(kelvin float64) float64{
	kelvin -= 273.15
	return kelvin
}

func celsiusToFahrenheit(celsius float64) float64{
	celsius = (celsius * 9.0/5.0) + 32.0
	return celsius
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println("kelvin:", kelvin, "celsius:", celsius, "fahrenheit", fahrenheit)
}
