package main

import (
	"fmt"
	"log"

	"github.com/2tvenom/golifx"
)

func main() {
	//Lookup all bulbs
	bulbs, err := golifx.LookupBulbs()
	if err != nil {
		log.Fatalf("Error retreving Bulbs: %v", err)
	}
	nameOfBulbAndPowerStatus(bulbs)

}

/*
	Description: loop through array of LIFX bulbs and get their name and location.
	Then print it out.
*/
func nameOfBulbAndPowerStatus(bulbs []*golifx.Bulb) {
	if !(len(bulbs) < 1) {
		for _, bulb := range bulbs {
			location, _ := bulb.GetLabel()
			powerState, _ := bulb.GetPowerState()
			fmt.Printf("Label: %s, Power: %v\n", location, powerState)
		}
	} else {
		fmt.Println("There are no bulbs available. Check your network connection or bulb connection. ")
	}
}
