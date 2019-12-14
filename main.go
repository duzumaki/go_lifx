package main

import (
	"fmt"

	"github.com/2tvenom/golifx"
)

func main() {
	//Lookup all bulbs
	bulbs, _ := golifx.LookupBulbs()
	nameOfBulbAndPowerStatus(bulbs)

}

/*
	Description: loop through array of LIFX bulbs and get their name and location.
	Then print it out.
*/
func nameOfBulbAndPowerStatus(bulbs []*golifx.Bulb) {
	for _, bulb := range bulbs {
		location, _ := bulb.GetLabel()
		powerState, _ := bulb.GetPowerState()
		fmt.Printf("Label: %s, Power: %v\n", location, powerState)
	}
}
