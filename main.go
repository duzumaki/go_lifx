package main

import (
	"fmt"

	"github.com/2tvenom/golifx"
)

func main() {
	//Lookup all bulbs
	bulbs, _ := golifx.LookupBulbs()

	//Turn if off
	// if !powerState {
	// 	bulbs[0].SetPowerState(true)
	// }

	// ticker := time.NewTicker(time.Second)
	// counter := 0

	// hsbk := &golifx.HSBK{
	// 	Hue:        2000,
	// 	Saturation: 13106,
	// 	Brightness: 65535,
	// 	Kelvin:     3200,
	// }
	// //Change color every second
	// // for _ = range ticker.C {
	// // 	bulbs[0].SetColorState(hsbk, 500)
	// // 	counter++
	// // 	hsbk.Hue += 5000
	// // 	if counter > 10 {
	// // 		ticker.Stop()
	// // 		break
	// // 	}
	// // }
	// //Turn off
	// bulbs[0].SetPowerState(false)
}

func nameOfBulbAndPowerStatus(bulbs []*golifx.Bulb) {
	for _, bulb := range bulbs {
		location, _ := bulb.GetLabel()
		powerState, _ := bulb.GetPowerState()
		fmt.Printf("Label: %s, Power: %v\n", location, powerState)
	}
}
