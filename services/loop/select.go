package loop

import (
	"fmt"
	"time"
)

// SelectRole1 select
/**
如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，否则的话，如果有default分支，则执行default分支语句，如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行

所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右

break关键字结束select

goto loop
break loop
continue loop

*/
func SelectRole1() {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	done := make(chan struct{}, 0)
	go func() {
		for i := 0; i < 1; i++ {
			ch2 <- i
		}

	}()
	go func() {
		for i := 0; i < 1; i++ {
			ch1 <- i
		}
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

Loop:
	for {
		select {
		case <-ch1:
			fmt.Println(1111)
		case <-ch2:
			fmt.Println(2)
		case <-done:
			fmt.Println("done")
			break Loop
		}
	}
}

// SelectBreak select break
func SelectBreak() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 3
	ch2 <- 5
	select {
	case <-ch1:

		fmt.Println("ch1 selected.")

		break

	case <-ch2:

		fmt.Println("ch2 selected.")
		fmt.Println("ch2 selected without break")
	}
}

// GotoLoopHelper goto loop
func GotoLoopHelper() {
	for a := 0; a < 5; a++ {
		fmt.Println(a)
		if a > 3 {
			goto Loop
		}
	}
Loop: //放在for后边. 放在for前面会一直执行。
	fmt.Println("test")
}

// BreakLoopHelper break loop
func BreakLoopHelper() {
Loop:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 1 {
				fmt.Println("a>1 continue loop")
				//continue Loop
				continue Loop
			}
		}
		if j > 1 {
			fmt.Println("j>1 break loop")
			break Loop
		}
	}
}
