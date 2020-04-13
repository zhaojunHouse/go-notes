### TODO
context
channel
defer
range
select
goto break continue
Panic recover
闭包closure


如何测试单个函数
rpc如何调试 curl



supervisor
log
网关
load config
HTTP server
RPC server
Tracing
PProf
interface
mysql connection pool
redis connection pool
kafka
grpc
锁问题
分布式锁
并发问题
分布式事务一致性问题。
git CICD
分布式
项目整体结构。

mac修改环境变量
git https ssh
sourceTree gitlab配置
git clone http://user:passwor@host

gopath goroot
go mod tidy


## 1. 声明/变量/赋值

### 声明关键字
	const
  	var
   	func
   	type
  
   
### 声明&初始化&赋值
    var 变量名 （类型） （= 初始值）  括号内可以省略一个
	//1
	var a int
	a = 1
	
	//2
	a := 1
	
	//3
	a := func(someFunc)
	
	//4
	medals := []string{"gold", "silver", "bronze"}
	
### 全局变量/局部变量/导出变量
	变量大写导出

  
### 指针/生命周期
	函数内部的局部变量函数退出时不一定回收
	引用或指针变量不可达，就表示可以回收

### 类型（type）
	type 类型名字 底层类型
	type Z int

### 包&文件&作用域


# 接口
    1 参数可以是接口类型
    2 接口声明时可内嵌其他接口
    3 实现接口：实现方法
    4 interface{}表示任何类型
    5 flag包
    6 内嵌，继承，重写
    7 接口值：动态类型/动态值
    8 断言
    用法：表示各种类型；接口实现；
    
# goroutine&channel
    声明，操作，关闭，阻塞
    当一个channel被关闭后，再向该channel发送数据将导致panic异常。
    当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。
    单向channel，缓存channel
    goroutine泄漏  runtime.NumGroutines
    匿名函数中的循环变量快照问题
    wg.add  done   wait
    select   channel=nil禁用case
    并发：防止太多并发。close channel，泄漏，阻塞死锁
    退出：close
 
 # 共享变量的并发   
    两个以上的groutine其中至少有一个写操作，就存在条件竞争。
    A1： 不写
    A2： 在一个groutine中操作变量
    A3： 互斥锁：不能重入
    
    读锁：内存同步 rlock
    写锁：lock
    
    runtime竞争条件检测
    
    
  #测试
    测试函数、基准测试(benchmark)函数、示例函数
    
   # 反射
    typeof valueof
    reflect.Value转原始类型
    获取类型底层类型
    遍历字段和方法
    修改字段的值
    动态调用方法