package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	isCall     bool
	isDone     bool
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
	// 認證
	tok := plurkAuth(&c)
	// 取得使用者資料
	opt := map[string]string{}
	opt["include_plurks"] = "false"
	ans, _ := callAPI(tok, "/APP/Profile/getOwnProfile", opt)
	plurker := plurkerObj{} // 使用者
	json.Unmarshal(ans, &plurker)
	printObjIndent(plurker)

	for true {
		//取得最近的噗
		opt = map[string]string{}
		opt["offset"] = time.Now().Format("2006-1-2T15:04:05") //現在時間
		opt["limit"] = "10"
		opt["minial_user"] = "true"
		ans, _ := callAPI(tok, "/APP/Timeline/getPlurks", opt)
		plurks := plurksObj{} // 抓回來的噗
		json.Unmarshal(ans, &plurks)
		isCall := false
		isDone := false // 是否結束
		for _, plurk := range plurks.Plurks {
			isCall = strings.Contains(plurk.ContentRaw, "匯率") // 有開村字串
			if isCall {
				// 取得回應
				opt = map[string]string{}
				opt["plurk_id"] = strconv.Itoa(plurk.PlurkID)
				opt["minimal_user"] = "true"
				ans, _ = callAPI(tok, "/APP/Responses/get", opt)
				responses := plurksObj{}
				json.Unmarshal(ans, &responses)
				for _, response := range responses.Responses { // 每個回應
					fmt.Println(response.ContentRaw)
					if !isDone {
						isDone, _ = regexp.MatchString("匯率", response.ContentRaw)
						fmt.Println(response.ContentRaw)
					}
				}
			}
		}
		break
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

func printJSONIndent(data []byte, indent string) {
	var jsi bytes.Buffer
	json.Indent(&jsi, []byte(data), "", indent)
	fmt.Printf("\n%s\n\n", jsi.Bytes())
}

func printObjIndent(data interface{}) {
	ans, _ := json.Marshal(data)
	printJSONIndent(ans, "  ")
}

//格式化輸出並換行
func printfln(format string, a ...interface{}) {
	fotmatSrt := fmt.Sprintf(format, a)
	fmt.Println(fotmatSrt)
}
