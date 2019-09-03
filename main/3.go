package main

import (
	"fmt"
	"reflect"
)

type User struct{
	Name string
	Age int
}

func (u User) Abs2(){
	print(1)
}

func (u *User) abs1()  {
	print(2)
}

func (u User) Print(prfix string){
	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}

func main() {
	u:= User{"张三",20}
	t:=reflect.TypeOf(u)
	fmt.Println(t)

	v:=reflect.ValueOf(u)
	fmt.Println(v)

	u1:=v.Interface().(User)
	fmt.Println(u1)

	fmt.Println(t.Kind())

	for i:=0;i<t.NumField();i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i:=0;i<t.NumMethod() ;i++  {
		fmt.Println(t.Method(i).Name)
	}

}

