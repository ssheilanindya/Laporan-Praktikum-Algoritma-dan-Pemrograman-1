package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Print("Masukkan suhu dalam Fahrenheit :")
	fmt.Scanln(&fahrenheit)
	celcius := (5.0 / 9.0) * (fahrenheit - 32)
	fmt.Printf("Suhu dari satuan fahrenheit ke celcius adalah: %.2f\n", celcius)
}
