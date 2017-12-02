package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type JSONTime time.Time

/*EsunBank ...
  玉山銀行類型
*/
type EsunBank struct {
	Bank
}

type dType struct {
	Rates     []esunRate
	QuoteTime JSONTime
}

type esunRate struct {
	Name           string
	CashBBoardRate float64
	URL            string `json:"Url"`
	Alt            string
	CashBonus      float64
	BBoardRate     float64
	CCY            string
	Key            string
	Bonus          float64
	Serial         int
	UpdateTime     JSONTime
	Title          string
	SBoardRate     float64
	CashSBoardRate float64
	Description    string
}

/*GetRate ...
實現玉山銀行取得匯率
*/
func (b *EsunBank) GetRate() (*[]BankRate, error) {
	var rate []BankRate
	datas := doGetEsunRate()
	for _, element := range datas {
		item := BankRate{
			Currancy: element.Key,
			CashSell: element.CashSBoardRate,
			CashBuy:  element.CashBBoardRate,
			SpotSell: element.SBoardRate,
			SpotBuy:  element.BBoardRate,
		}
		//check has to Only Currancy
		if len(b.Currancy) > 0 && b.Currancy != item.Currancy {
			continue
		}
		rate = append(rate, item)
	}
	return &rate, nil
}

func doGetEsunRate() []esunRate {
	url := "https://www.esunbank.com.tw/bank/Layouts/esunbank/Deposit/DpService.aspx/GetForeignExchageRate"
	t := time.Now()
	payload := strings.NewReader("{day:'" + t.Format("2006-01-02") + "',time:'" + t.Format("15:04:05") + "'}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	req.Header.Add("Referer", "https://www.esunbank.com.tw/bank/personal/deposit/rate/forex/foreign-exchange-rates")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// interface 可以接收任何類型的值
	u := map[string]interface{}{}
	err := json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println(err.Error)
	}
	str := []byte(u["d"].(string))
	r := dType{}
	err = json.Unmarshal(str, &r)
	return r.Rates
}
