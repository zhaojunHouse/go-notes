package uniqueID

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

func CreateUUID(){
	// 创建
	u1, err := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}


func IDWroker(){
     payOrderID := "12145242"
     timestamp := strconv.FormatInt(time.Now().UnixNano(),10)
     fmt.Println(payOrderID)
     fmt.Println(timestamp)
     fmt.Println(payOrderID+timestamp)
     fmt.Println(strings.Join([]string{payOrderID,timestamp},"_"))
}