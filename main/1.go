package main

import "fmt"

/**
声明，初始化，赋值
*/

//声明  type var const func
type a int

var b int

const c int = 1

func Add(a int, b int) int {
	return a + b
}

//变量   var 变量名  (变量类型)  (= 初始值)
var a1 = 1
var a2 int
var a3 int = 3

var a4 [3]int
var a5 *[2]string

type a6 struct {
	a7 bool
}


//pointer
func incr(p *int) *int{
	*p++
	return p
}

func main() {
	a := a6{}
	p := &a

	po := 0
	for i:=0;i<10;i++ {
		incr(&po)
		fmt.Println(po)
	}


	(*p).a7 = true
	fmt.Println(a.a7)
}
