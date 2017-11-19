package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	plurgo "github.com/clsung/plurgo/plurkgo"
	"github.com/garyburd/go-oauth/oauth"
)

var (
	c          string
	d          bool
	h          bool
	l          int
	opt        map[string]string
	errc       int
	taiwanBank *TaiwanBank
	esunBank   *EsunBank
)

func init() {
	flag.StringVar(&c, "c", "config.json", "載入設定檔")
	flag.BoolVar(&d, "d", false, "刪除所有噗")
	flag.BoolVar(&h, "h", false, "說明")
	flag.IntVar(&l, "l", -1, "紀錄")
	flag.Usage = usage
	taiwanBank = &TaiwanBank{Bank: Bank{Name: "Bank of Taiwan", CN: "台灣銀行", URL: "http://rate.bot.com.tw/xrt/flcsv/0/day"}}
	esunBank = &EsunBank{Bank: Bank{Name: "Esun Bank", CN: "玉山銀行"}}
}

func main() {
	token := plurkAuth(&c)
	if token == nil {
		fmt.Print("err")
	}
	// taiwanBank.Currancy = "USD"
	taiwanRates, err := taiwanBank.GetRate()
	if err != nil {
		printfln("%v", err)
	}
	rateJson, _ := json.Marshal(taiwanRates)
	printfln("%v", string(rateJson))

}

//v顯示指令列表
func usage() {
	fmt.Printf("\n%s\n", "Options:")
	flag.PrintDefaults()
	fmt.Println()
}

//plurk OAuth2.0 授權
func plurkAuth(credPath *string) *oauth.Credentials {
	flag.Parse()
	plurkOAuth, err := plurgo.ReadCredentials(*credPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	toekn, authorized, err := plurgo.GetAccessToken(plurkOAuth)
	if authorized {
		bytes, err := json.MarshalIndent(plurkOAuth, "", " ")
		if err != nil {
			log.Fatalf("failed to store credential :%+v", err)
		}
		err = ioutil.WriteFile(*credPath, bytes, 0700)
		if err != nil {
			log.Fatalf("failed to write credential :%+v", err)
		}
	}
	return toekn
}

//呼叫api
func callAPI(token *oauth.Credentials, api string, opt map[string]string) ([]byte, error) {
	ans, e := plurgo.CallAPI(token, api, opt)

	if e != nil {
		errc++
		log.Fatal(e)
	} else {
		errc = 0
	}
	return ans, e
}

//格式化輸出並換行
func printfln(format string, a ...interface{}) {
	fotmatSrt := fmt.Sprintf(format, a)
	fmt.Println(fotmatSrt)
}
