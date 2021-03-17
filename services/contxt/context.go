package contxt

import (
	"context"
	"fmt"
	"time"
)

// CancelContext cancel context
func CancelContext() {
	gen := func(ctx context.Context) <-chan int {
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// ContextWithDeadline context超时
func ContextWithDeadline() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
Loop:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("sleep 1 second")
		case <-ctx.Done():
			fmt.Println("ctx deadline")
			fmt.Println(ctx.Err())
			break Loop
		}
	}
}

// ContextWithTimeout context超时
func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

Loop:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("sleep 1 second")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			fmt.Println("ctx time out")
			break Loop
		}
	}

}

// ContextWithValue context设置值
func ContextWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
