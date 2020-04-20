package uniqueid

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// IDWroker create id
func IDWroker() {
	payOrderID := "12145242"
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(payOrderID)
	fmt.Println(timestamp)
	fmt.Println(payOrderID + timestamp)
	fmt.Println(strings.Join([]string{payOrderID, timestamp}, "_"))
}
