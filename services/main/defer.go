package main

import "fmt"

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}


func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main(){
	var str string
	str = "hellowangxiaochao"
	fmt.Printf("%+v\n",str[:5])
}

git config --global url."https://gitlab.com/".insteadOf "git@gitlab.com:"