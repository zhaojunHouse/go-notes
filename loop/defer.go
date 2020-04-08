package loop

import "fmt"

/**
 规则一:defer表达式中变量的值在defer表达式被定义时就已经明确
 规则二:defer表达式的调用顺序是按照先进后出的方式
 规则三:defer表达式中可以修改函数中的命名返回值
 规则四:defer在panic之后不执行

defer 中先执行的s.Add(1).Add(4).Add(100)，然后执行s.Add(11)， s.Add(12)，s.Add(13)，延迟执行的test函数，可以看到defer延迟执行的是最后的一个函数
*/

/**
使用场景：
资源收回
panic异常的捕获
修改函数命名返回值
*/

// panic异常的捕获
func DeferPanic() {
	f()
	fmt.Println("Returned normally from f.")
}

// 规则三:defer表达式中可以修改函数中的命名返回值
func DeferUpdateFuncResult() {
	fmt.Println(test())
}

/**
start 7
end 8
8
*/
func test() (res int) {
	res = 1
	defer func() {
		fmt.Println("start", res)
		res++
		fmt.Println("end", res)
	}()
	return 7
}

/**
规则一:defer表达式中变量的值在defer表达式被定义时就已经明确
规则二:defer表达式的调用顺序是按照先进后出的方式
*/
func DeferHelper() {
	//i := 0
	//defer fmt.Println(i)
	//i++
	//defer fmt.Println(i)

	//fmt.Println(f1())    // 返回1
	//fmt.Println(f2())      // 返回5
	//fmt.Println(f3())     // 返回10
	fmt.Println(f4()) // 返回2
	return
}

// defer可以修改函数命名返回值。
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// defer可以修改函数命名返回值。r = t
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// defer可以修改函数命名返回值。
func f3() (t int) {
	t = 5
	defer func() {
		t = t + 5
	}()
	return t
}

// defer带参数值的闭包。
func f4() (r int) {
	defer func(r int) {
		fmt.Println("defer--r", r)
		r = r + 5
		fmt.Println("defer--r", r)
	}(r)
	return 2
}

// defer 闭包， 闭包传递的是指针。
func DeferBibao() {
	//fmt.Println(a())
	fmt.Println(b())
}

// 如果你在定义defer的时候,就要将defer后面的函数参数等入栈,等到ruturn之前的时候出栈执行,a中是将i的拷贝直接入栈,b中通过一个闭包调用,实际上将i的指针传递给闭包,闭包读取值拷贝给add.
func a() int {
	var i int
	defer add(i) //这里虽然defer是在return之前执行,但是在定义的时候,
	// 已经将defer要执行的函数压入栈,所以传递给add的是var i int的i值.
	/*defer func(){
	    add(i)
	}() */
	i += 100
	return i //return  0
}

func b() int {
	var i int
	defer func() {
		add(i)
	}()
	i += 100
	return i //return  0
}

func add(i int) {
	i += 1
	fmt.Println("add----", i)

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")
}

func g() {
	panic("ERROR")
}
