package main

import "fmt"

const boiling = 373

func main() {

	var tempK = boiling
	var tempC = (tempK - 273)

	fmt.Println("Temperatura de ebolição em Kelvin é: ", tempK)
	fmt.Println("Temperatura de ebolição em Celcius é: ", tempC)
}
