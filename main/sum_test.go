package main

import "testing"

func TestSum(t *testing.T) {
	s := Sum(3,5)
	if s != 8 {
		t.Error("sum func failed")
	}
	t.Log("sum result is :",s)
}


func Test_stuct_nil (t *testing.T) {

}