package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Printf("%d left", fuel)
}

func calculateFuel(planet string) (int) {
	var fuel int
	if planet == "Mercury" {
		fuel = 500000
	} else if planet == "Venus" {
		fuel = 300000
	} else if planet == "Mars" {
		fuel = 700000
	}
	return fuel
}

func greetPlanet (planet string) {
	fmt.Print("Welcome to planet ", planet)
}

func cantFly() {
	fmt.Print("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) (int){
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	fuel := 1000000
	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)  
}