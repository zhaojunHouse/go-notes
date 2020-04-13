package loop

import (
	"fmt"
	"time"
)

func PanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("func begin")
	a := []string{"a", "b"}
	fmt.Println(a[3])       // 越界访问，肯定出现异常
	panic("bug")            // 上面已经出现异常了,所以肯定走不到这里了。
	fmt.Println("func end") // 不会运行的.
	time.Sleep(1 * time.Second)
}
