package loop

import (
	"fmt"
	"time"
)

// 坑1 ：for range 中使用闭包
func ClosureBug1(){
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(2*time.Second)
}


// 坑2：函数列表使用不当
func ClosureBug2(){
	for _, f := range funcMap() {
		f()
	}
}

func funcMap() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		x := i                  //复制变量
		s = append(s, func() {
			fmt.Println(&x, x)
		})
	}

	return s
}

// 坑3 ：延迟调用
func ClosureDefer(){
	x, y := 1, 2

	defer func(a int) {
		fmt.Printf("x:%d,y:%d\n", a, y)  // y 为闭包引用
	}(x)      // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)
}