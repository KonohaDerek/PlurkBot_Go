package main

import (
	"fmt"
	"strconv"
)

//Bank 銀行類別基底
type Bank struct {
	Name     string
	CN       string
	URL      string
	Currancy string
}

/*EsunBank ...
  玉山銀行類型
*/
type EsunBank struct {
	Bank
}

/*BankRate ...
  銀行匯率基底
*/
type BankRate struct {
	Currancy string
	CashSell float64
	CashBuy  float64
	SpotSell float64
	SpotBuy  float64
}

/*Rate ...
  匯率接口
*/
type Rate interface {
	GetRate() []BankRate
}

/*GetRate ...
實現玉山銀行取得匯率
*/
func (b *EsunBank) GetRate() (*[]BankRate, error) {
	var rate []BankRate
	rate = append(rate, BankRate{"JPY", 0.267, 0.267, 0.268, 0.269})
	if b == nil {
		return nil, fmt.Errorf("err")
	}

	if b.Currancy == "" {
		return &rate, nil
	}
	return &rate, nil
}

func convertToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("%v : %v not float nu,ber", err, str)
		return 0
	}
	return value
}
