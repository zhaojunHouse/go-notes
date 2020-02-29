package main

import (
	"fmt"
	"reflect"
)

type ItemOption struct {
	itemVisibleOption int32 `json:"item_visible_option"`
}
type Item struct {
	ID     int64       `json:"id"`
	stock  int32       `json:"stock"`
	Option *ItemOption `json:"option"`
}

func main() {
	item := &Item{
		ID:    1,
		stock: 2,
		Option: &ItemOption{itemVisibleOption:int32(22)},
	}
	t := reflect.TypeOf(*item)
	//TODO  结构体中是否包含某个属性
	if _, ok := t.FieldByName("Option"); ok {
		fmt.Println("exist")
		//TODO  指针类型必须判空
		if item.Option != nil {
			fmt.Println(item.Option.itemVisibleOption)
		}
	}
	if _, ok := t.FieldByName("Name"); !ok {
		fmt.Println(" not exist")
	}
	fmt.Println(item.ID)
	fmt.Println(item.stock)
	fmt.Println(item.Option)

}
