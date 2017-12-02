package main

import (
	"encoding/json"
	"testing"
)

func Test_doGetEsunRate(t *testing.T) {
	doGetEsunRate()
}

func Test_GetEsunRate(t *testing.T) {
	esunBank = &EsunBank{Bank: Bank{Name: "Esun Bank", CN: "玉山銀行"}}
	esunRates, err := esunBank.GetRate()
	if err != nil {
		printfln("%v", err)
	}
	rateJson, _ := json.Marshal(esunRates)
	printfln("Esun Rates : %v", string(rateJson))
}
