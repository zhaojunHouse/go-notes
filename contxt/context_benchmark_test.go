package contxt

import (
	"fmt"
	"testing"
)

/**
参考：http://c.biancheng.net/view/124.html
*/

//  go test -v -bench=. benchmark_test.go
func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

// go test -v -bench=Alloc -benchmem benchmark_test.go
// go test -v -bench=Alloc -benchtime=5s benchmark_test.go
/**
    Benchmark_Alloc-4        9052664               131 ns/op              16 B/op          2 allocs/op
	4 ： 执行groutinue的数量
	131 ns/op    每次执行耗时131ns
	2 allocs/op  每次执行有2次分配
	16 B/op      每次执行有16字节分配

*/
func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}
