package time

import (
	"fmt"
	"time"
)

func timeAdd() {
	now := time.Now()
	fmt.Println(now)
	yesterday := now.Add(-time.Duration(time.Second * 24 * 3600))
	fmt.Println(yesterday)
	fmt.Println(time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0, time.Local))
	tomorrow := now.Add(time.Duration(time.Second * 24 * 3600))
	fmt.Println(tomorrow)

}
