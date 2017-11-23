package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*EsunBank ...
  玉山銀行類型
*/
type EsunBank struct {
	Bank
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

func doGetEsunRate() {
	url := "https://www.esunbank.com.tw/bank/Layouts/esunbank/Deposit/DpService.aspx/GetForeignExchageRate"
	payload := strings.NewReader("{day:'2017-11-19',time:'22:59:48'}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
