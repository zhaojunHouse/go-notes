package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

type User struct {
	Name string
	Age  int
}

func (u User) Abs2() {
	print(1)
}

func (u *User) abs1() {
	print(2)
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func Sum(a int, b int) int {
	return a + b
}

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func Gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx canceled")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers
	//
	//for n := range Gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		fmt.Println("break from loop")
	//		break
	//	}
	//}



	//d := time.Now().Add(50000 * time.Millisecond)
	//ctx, cancel := context.WithDeadline(context.Background(), d)
	//
	//// Even though ctx will be expired, it is good practice to call its
	//// cancelation function in any case. Failure to do so may keep the
	//// context and its parent alive longer than necessary.
	//defer cancel()
	//
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}



	//type favContextKey string
	//
	//f := func(ctx context.Context, k favContextKey) {
	//	if v := ctx.Value(k); v != nil {
	//		fmt.Println("found value:", v)
	//		return
	//	}
	//	fmt.Println("key not found:", k)
	//}
	//
	//k := favContextKey("language")
	//ctx := context.WithValue(context.Background(), k, "Go")
	//
	//f(ctx, k)
	//f(ctx, favContextKey("color"))


	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
