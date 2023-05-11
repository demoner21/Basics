package main

import "fmt"

const Ebuli float64 = 212.0

func main() {

	var TempF = Ebuli
	var TempC = (TempF - 32) * 5 / 9

	fmt.Println("Ponto de Ebolição da àgua em ºF: ", TempF)
	fmt.Println("Ponto de Ebolição da àgua em ºC: ", TempC)
}
