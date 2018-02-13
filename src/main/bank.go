package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Bank 銀行類別基底
type Bank struct {
	Name     string
	CN       string
	URL      string
	Currancy string
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
	RateTime time.Time
	BankInfo Bank
}

/*Rate ...
  匯率接口
*/
type Rate interface {
	GetRate() []BankRate
}

/*
	轉換字串>浮點樹
*/
func convertToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("%v : %v not float nu,ber", err, str)
		return 0
	}
	return value
}

/*
	格式化輸出訊息
*/
func BankRateFotmat(rate BankRate) string {
	strs := []string{}
	//填入銀行名稱
	strs = append(strs, fmt.Sprintf("%s匯率", rate.BankInfo.CN))
	//加入時間
	if !rate.RateTime.IsZero() {
		strs = append(strs, fmt.Sprintf("時間 : %s", rate.RateTime.Local().Format("2006-01-02 15:04:05")))
	}

	//加入現金賣價
	strs = append(strs, fmt.Sprintf("現金(賣) : %f", rate.CashSell))
	//加入即期賣價
	strs = append(strs, fmt.Sprintf("即期(賣) : %f", rate.SpotSell))
	//加入現金買價
	strs = append(strs, fmt.Sprintf("現金(買) : %f", rate.CashBuy))
	//加入即期買價
	strs = append(strs, fmt.Sprintf("即期(買) : %f", rate.SpotBuy))
	//輸出
	return strings.Join(strs, "\n")
}
