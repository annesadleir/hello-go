package main

import (
	"fmt"
	"time"
)

const SECONDS_PER_HOUR int64 = 3600

func main() {
	now := time.Now()
	nowUnix := now.Unix()
	fmt.Println(now)
	fmt.Println(nowUnix)
	// ascertain the nearest hour to now
	lastHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	fmt.Println(lastHour)
	lastHourUnix := lastHour.Unix()
	fmt.Println(lastHourUnix)
	fmt.Println(nowUnix - (nowUnix % SECONDS_PER_HOUR))
}
