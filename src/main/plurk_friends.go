package main

import (
	"time"

	"github.com/garyburd/go-oauth/oauth"
)

func AutoAddFriends(token *oauth.Credentials) {
	//取得最近的噗
	opt = map[string]string{}
	for true {
		callAPI(token, "/APP/Alerts/addAllAsFriends", nil)
		time.Sleep(30 * time.Second)
	}
}
