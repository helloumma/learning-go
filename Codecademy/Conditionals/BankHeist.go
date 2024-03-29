package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn:=true
	eludedGuards:=rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn=false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault:=rand.Intn(100)
	fmt.Println(isHeistOn)

	if isHeistOn && openedVault >=70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		fmt.Println("Vault can't be opened.")
	} else {
		fmt.Println("what's the combo to this lock again??")
	}

	leftSafely:=rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			fmt.Print("Turns out the vault doors don't open from the inside...")
		case 2 :
			fmt.Print("What an earth is happening?")
		case 3: 
			fmt.Print("ugh...?")
		default:
			fmt.Print("Start the getaway car!")
		}
		amtStole:= 1000+rand.Intn(1000000)
		fmt.Print(amtStole)
		
	}
}
