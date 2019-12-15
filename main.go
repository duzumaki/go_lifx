package main

import (
	"fmt"
	"log"
	"os"

	"github.com/2tvenom/golifx"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// // Twitter client
	client := twitter.NewClient(httpClient)
	//run get command to get json and check the latest tweet that mentions a color and change it to that color.
	// User Show
	t := true
	f := false
	tweets, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{Count: 3, IncludeEntities: &f, TrimUser: &t})
	//https://lan.developer.lifx.com/docs/building-a-lifx-packet
	// body, err := ioutil.ReadAll(.Body)
	tweetColor := tweets[1].Text
	fmt.Println(tweetColor)

	hsbk := &golifx.HSBK{
		Hue:        63519,
		Saturation: 45000,
		Brightness: 65535,
		Kelvin:     0,
	}

	if tweetColor == "@uzumakithegod pink" {
		for i := 0; i < len(bulbs); i++ {
			bulbs[i].SetColorState(hsbk, 100)
		}
	}

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
