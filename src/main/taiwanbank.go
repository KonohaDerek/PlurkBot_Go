package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*TaiwanBank ...
  台灣銀行類型
*/
type TaiwanBank struct {
	Bank
}

/*GetRate ...
實現台灣銀行取得匯率
*/
func (b *TaiwanBank) GetRate() (*[]BankRate, error) {
	var rate []BankRate
	if b == nil {
		return nil, fmt.Errorf("Not Have Taiwan Information")
	}
	_, lines, err := readCSVFromURL("http://rate.bot.com.tw/xrt/flcsv/0/day")

	if err != nil {
		log.Fatalf("%+v", err)
	}

	for index, line := range lines {
		if index == 0 {
			continue
		}
		//Split the CSV Separated ',' to get the each colum
		splitStr := strings.Split(line[0], ",")
		item := BankRate{
			Currancy: splitStr[0],
			CashSell: convertToFloat(splitStr[13]),
			CashBuy:  convertToFloat(splitStr[3]),
			SpotSell: convertToFloat(splitStr[12]),
			SpotBuy:  convertToFloat(splitStr[2]),
		}
		//check has to Only Currancy
		if len(b.Currancy) > 0 && b.Currancy != item.Currancy {
			continue
		}
		rate = append(rate, item)
	}
	return &rate, nil
}

/*
	從URL中直接取的CSV檔案
*/
func readCSVFromURL(url string) ([]string, [][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	//first line of CSV file that defines our json object properties
	properties := strings.Split(lines[0][0], ",")
	return properties, lines, nil
}
