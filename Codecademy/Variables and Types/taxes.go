package main

import "fmt"

func main() {
	coolSneakers := 65.99
	niceNecklace := 45.50
	var taxCalculation float64
	taxCalculation += coolSneakers
	taxCalculation += niceNecklace
	taxCalculation *= 0.08875
	fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)
}