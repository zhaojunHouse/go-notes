package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	todayZeroClock := t.Unix() - 8*3600
	todayTwentyFourClock := t.Unix() - 8*3600 + 24*3600
	fmt.Println("today_zero_point", todayZeroClock)
	fmt.Println("today_zero_point", todayTwentyFourClock)
}
