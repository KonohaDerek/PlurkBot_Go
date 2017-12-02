package main

import (
	"encoding/json"
	"testing"
)

func Test_GetTaiwanRate(t *testing.T) {
	// taiwanBank.Currancy = "USD"
	taiwanRates, err := taiwanBank.GetRate()
	if err != nil {
		printfln("%v", err)
	}
	rateJson, _ := json.Marshal(taiwanRates)
	printfln("%v", string(rateJson))
}
