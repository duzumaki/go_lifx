package main

import (
	"fmt"

	"github.com/2tvenom/golifx"
)

func main() {
	//Lookup all bulbs
	bulbs, _ := golifx.LookupBulbs()
	//fmt.Println(bulbs)
	//Get label
	for _, bulb := range bulbs {
		location, _ := bulb.GetLabel()
		fmt.Printf("Label: %s\n", location)
	}

	//Get power state
	//powerState, _ := bulbs[0].GetPowerState()

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
