package goroutine

import (
	"fmt"
	"time"
)

/**
参考
https://blog.csdn.net/cgl1079743846/article/details/90691146
http://litang.me/post/golang-channel/
*/

// channel特性
// 从一个nil channel接收数据会一直阻塞
// 如果channel没有数据，只要channel未关闭，从channel中接收数据会一直阻塞直到有数据为止
// 无缓冲channel，发送数据和接收数据同时发生。如果没有receiver接收数据(<- chan)，则sender发送数据会一直阻塞；如果没有sender发送数据(chan <-)，则receiver接收数据会一直阻塞
// 有缓冲channel，当队列容器未满的时候sender不会阻塞，当队列容量满的时候sender会阻塞，当队列容量为空的时候receiver会阻塞
// 关闭nil channel会panic (panic: close of nil channel)
// 重复关闭channel会panic (panic: close of closed channel)
// 向已经关闭channel发送数据会panic (panic: send on closed channel)
// 从已经关闭的channel读取数据不会阻塞，如果channel为空读到对应类型默认值，例如int默认值是0，指针默认值为nil
// channel是一个FIFO队列，发送和接收的顺序是一致的
// 可以使用val, ok := <-ch 方式来判断channel是否关闭，ok为true表示channel未关闭，ok为false表示channel关闭
// 支持多个goroutine同时消费一个Channel，可用于并发处理场景

// Channel最常见用法主要是以下6种
//
//goroutine之间同步通信
//range迭代
//select操作
//timeout超时处理
//timer计时器
//ticker定时器

/**
futures / promises
一对一通知
*/

// ChannelWaitingCompleted channel用法1 ： 等待groutine完成
func ChannelWaitingCompleted() {
	println("start main")
	ch := make(chan bool)
	defer close(ch)
	go func() {
		println("come into goroutine")
		ch <- true
	}()

	println("do something else")
	<-ch

	println("end main")
}

// ChannelSum channel 用法2 ：多个groutine协同
func ChannelSum() {
	println("start main")
	ch := make(chan int)
	defer close(ch)

	var result int
	go func() {
		println("come into goroutine1")
		r := 1
		ch <- r
	}()

	go func() {
		println("come into goroutine2")
		r := 2
		ch <- r
	}()

	go func() {
		println("come into goroutine3")
		ch <- 3
	}()

	for i := 0; i < 3; i++ {
		result += <-ch
	}

	println("result is:", result)
	println("end main")
}

// ChannelSelect channel用法3： select
func ChannelSelect() {
	println("start main")
	cond1 := make(chan int)
	cond2 := make(chan uint64)

	go func() {
		for i := 0; ; i++ {
			cond1 <- i
		}
	}()

	go func() {
		var i uint64
		for ; ; i++ {
			cond2 <- i
		}
	}()

	endCond := false
	for endCond != true {
		select {
		case a := <-cond1:
			if a > 99 {
				println("end with cond1", "a=", a)
				endCond = true
			}
		case b := <-cond2:
			if b == 100 {
				println("end with cond2", "b=", b)
				endCond = true
			}
		case <-time.After(time.Microsecond):
			println("end with timeout")
			endCond = true
		}
	}

	println("end main")
}

// ChannelBuffer 带缓冲的channel
func ChannelBuffer() {
	println("start main")
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	println("end main")
}

// ChannelNoBuffer channel no buffer
func ChannelNoBuffer() {
	var ch = make(chan int)
	go func() {
		ch <- 1
		println("sender")
	}()
	println(<-ch)
}

// ChannelClose 广播
func ChannelClose() {
	N := 10
	exit := make(chan struct{})
	done := make(chan struct{}, N)
	// start N worker goroutines
	for i := 0; i < N; i++ {
		go func(n int) {
			for {
				select {
				// wait for exit signal
				case <-exit:
					fmt.Printf("worker goroutine #%d exit\n", n)
					done <- struct{}{}
					return
				case <-time.After(time.Second):
					fmt.Printf("worker goroutine #%d is working...\n", n)
				}
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
	// broadcast exit signal
	close(exit)
	// wait for all worker goroutines exit
	for i := 0; i < N; i++ {
		<-done
	}
	fmt.Println("main goroutine exit")
}
