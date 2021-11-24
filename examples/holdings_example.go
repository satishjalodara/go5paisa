// Example to fetch all holdings for a user
package main

import (
	"fmt"
	"github.com/5paisa/go5paisa"
)

func main() {

	conf := &go5paisa.Config{
		AppName:       "5P53076904",
		AppSource:     "4223",
		UserID:        "5pcntfcOscA",
		Password:      "qv5GtT63l80",
		UserKey:       "UnSC1c4nrRhMmJMahoZmlN8jXo6Rvruu",
		EncryptionKey: "wPTRvBs6xLrL5cJFmwibXSWhFHaGTdSc",
	}
	appConfig := go5paisa.Init(conf)
	client, err := go5paisa.Login(appConfig, "8vQRCYZWhIrQ8RREx7GKHolrWe6jqlgZb3zjvxVHpJI=", "rPLQ/ENmoy4GVyMXt8x3FA==", "8hgQM+jqMESAV1nxOoN")
	if err != nil {
		panic(err)
	}
	holdings, err := client.GetHoldings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", holdings)
}
