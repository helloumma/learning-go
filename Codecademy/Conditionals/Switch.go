package main

import "fmt"

func main() {
	name := "H. J. Simp"

	switch name {
	case "butch":
		fmt.Println("Head to Robbers Roost.")
	case "bonnie":
			fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}
}
