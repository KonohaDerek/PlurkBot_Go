package main

import (
	"time"
)

type JSONTime struct {
	time.Time
}

const CustomTimeFormat = "2006-01-02T15:04:05"

func (ct *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(CustomTimeFormat, string(b))
	return
}

func (ct *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(CustomTimeFormat)), nil
}
