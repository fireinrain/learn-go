package chapter4

import (
	"flag"
	"fmt"
	"strings"
)

const name = ":dsds"

func FunctionPlay() {

	//函数
	//Go 语言支持多返回值，多返回值能方便地获得函数执行后的多个返回参数，
	// Go 语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误。
	println(add(12, 3))
	fmt.Println(add(12, 34))
	fmt.Println(name)

	a, b := typedTwoValues()
	fmt.Println(a, b)

	fmt.Println("-----------")
	values, b2 := namedRetValues()
	fmt.Println(values, b2)

	// 将返回值作为打印参数
	fmt.Println(resolveTime(1000))
	// 只获取消息和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)
	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)
	fmt.Println(aaa)
	fmt.Println("-----------")

	//Go 语言中传入和返回参数在调用和返回时都使用值传递，
	// 这里需要注意的是指针、切片和 map 等引用型对象指向的内容在参数传递中不会发生复制，
	// 而是将指针进行复制，类似于创建一次引用。

	testPassByValue()

	fmt.Println("-----------")
	//函数变量
	var f func()
	f = fire
	f()
	//函数的链式调用
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	// 处理函数链,一个函数切片
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	//处理字符串
	StringProcess(list, chain)
	for _, value := range list {
		fmt.Println(value)
	}
	fmt.Println("----------")

	//匿名函数
	func(str string) {
		fmt.Println(strings.ToUpper(str))
	}("weizhuang")

	//匿名函数实现回调
	visit([]int{1, 2, 3}, func(i int) {
		fmt.Println(i, "abc")
	})

	var skillParam = flag.String("skill", "", "skill to perform")
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

	fmt.Println("---------------")

	//实现接口

	//声明接口变量
	var invoker Invoker

	mystruct := new(Struct)
	invoker = mystruct
	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function------", v)
	})

	invoker.Call("hello")

	//interface{} 表示任意类型

	//闭包
	//闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量。因此，简单的说：
	//函数 + 引用环境 = 闭包

	str := "hello world"
	foo := func() {
		str = "hello dude"
	}
	foo()
	fmt.Println("str value: ", str)
	fmt.Println("----------")

	//测试
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)

	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)

	// 累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
}

//闭包的记忆效应
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}

}

//实现接口
// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

//回调
func visit(list []int, f func(i int)) {
	for _, v := range list {
		f(v)
	}
}

//链式调用
//链式处理器是一种常见的编程设计。
// Netty 是使用 Java 语言编写的一款异步事件驱动的网络应用程序框架，
// 支持快速开发可维护的高性能的面向协议的服务器和客户端，Netty 中就有类似的链式处理器的设计。
//
//Netty 可以使用类似的处理链对封包进行收发编码及处理。
// Netty 的开发者可以分为 3 种：第一种是 Netty 底层开发者，
// 第二种是每个处理环节的开发者，第三种是业务实现者，在实际开发环节中，后两种开发者往往是同一批开发者。
// 链式处理的开发思想将数据和操作拆分、解耦，让开发者可以根据自己的技术优势和需求，进行系统开发，
// 同时将自己的开发成果共享给其他的开发者。
func StringProcess(list []string, chain []func(string2 string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func fire() {
	fmt.Println("fire")
}

type Data struct {
	complax []int

	instance InnerData

	ptr *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFun ptr: %p\n", &inFunc)
	return inFunc
}

func testPassByValue() {
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			a: 5,
		},
		ptr: &InnerData{
			a: 1,
		},
	}

	// 输入结构的成员情况
	//使用格式化的%+v动词输出 in 变量的详细结构
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)

	//所有的 Data 结构的指针地址发生了变化，意味着所有的结构都是一块新的内存，无论是将 Data 结构传入函数内部，还是通过函数返回值传回 Data 都会发生复制行为。
	//所有的 Data 结构中的成员值都没有发生变化，原样传递，意味着所有参数都是值传递。
	//Data 结构的 ptr 成员在传递过程中保持一致，表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分。

}

func add(a, b int) int {
	return a + b
}

func typedTwoValues() (int, int) {
	return 1, 2
}

func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

//秒转化为时间
const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

var aaa int = 12

func resolveTime(second int) (day int, hour int, minute int) {
	aaa = 13
	day = second / SecondsPerDay
	hour = second / SecondsPerHour
	minute = second / SecondsPerMinute
	return
}
