package main

import "fmt"

func main() {
	star := "Polaris"
	
	starAddress := &star
	
	fmt.Println("The address of star is", starAddress)
}
