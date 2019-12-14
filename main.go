package main

import (
	"fmt"
	"log"

	"github.com/2tvenom/golifx"
)

// NameAndPower ...
type NameAndPower struct {
	Name       []string
	PowerState []bool
}

func main() {
	//Lookup all bulbs
	bulbs, err := golifx.LookupBulbs()
	if err != nil {
		log.Fatalf("Error retreving Bulbs: %v", err)
	}
	nameOfBulbAndPowerStatus(bulbs, &NameAndPower{})

	// config := oauth1.NewConfig("consumerKey", "consumerSecret")
	// token := oauth1.NewToken("accessToken", "accessSecret")
	// httpClient := config.Client(oauth1.NoContext, token)

	// // Twitter client
	// client := twitter.NewClient(httpClient)

}

/*
	Description: loop through array of LIFX bulbs and get their name and location.
	Then print it out.
*/
func nameOfBulbAndPowerStatus(bulbs []*golifx.Bulb, info *NameAndPower) {
	if !(len(bulbs) < 1) {
		for _, bulb := range bulbs {
			location, _ := bulb.GetLabel()
			info.Name = append(info.Name, location)
			powerState, _ := bulb.GetPowerState()
			info.PowerState = append(info.PowerState, powerState)
			//fmt.Printf("Label: %s, Power: %v\n", location, powerState)
		}
		fmt.Println(info.Name[0], info.PowerState[0])
	} else {
		fmt.Println("There are no bulbs available. Check your network connection or bulb connection.")
	}
}
