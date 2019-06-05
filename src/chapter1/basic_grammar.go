package chapter1

import (
	"fmt"
	"math"
)

func BasicGrammar() {

	//声明变量 var 变量 类型
	//var a int
	//var b string
	//var c []float32
	//var d func() bool
	//var e struct{
	//	x int
	//}

	//a = 13
	//b = "a"

	//批量声明
	//var(
	//	a int
	//	b string
	//	c [] float32
	//	d func() bool
	//	e struct{
	//		x int
	//	  }
	//)

	//变量初始化
	/*
		Go 语言在声明变量时，自动对变量对应的内存区域进行初始化操作。每个变量会初始化其类型的默认值，例如：
		整型和浮点型变量的默认值为 0。
		字符串变量的默认值为空字符串。
		布尔型变量默认为 bool。
		切片、函数、指针变量的默认为 nil。
	*/
	//var hp int = 100;
	//fmt.Println("血量：",hp)

	//编译器推到
	var hp = 100
	fmt.Println(hp)

	//多变量声明初始化
	hpp := 100
	fmt.Println(hpp)

	// 多变量同时赋值
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)

	aa := 100
	bb := 200

	aa, bb = bb, aa
	fmt.Println(aa, bb)

	fmt.Println("--------------------------")
	var aaa int = 100
	var bbb int = 200
	aaa = aaa ^ bbb
	bbb = bbb ^ aaa
	aaa = aaa ^ bbb
	fmt.Println(a, b)

	// 匿名变量
	// 匿名变量不占用命名空间，不会分配内存。
	// 匿名变量与匿名变量之间也不会因为多次声明而无法使用
	a, _ = GetData()
	_, b = GetData()
	fmt.Println(a, b)
	fmt.Println("--------------------------")

	//整数类型
	//var e int = 12
	//var ee uint = -12
	//
	////浮点类型
	//var f  float32 = 2.3
	//var ff float64 = 2.33
	//
	//fmt.Println("%f\n",math.Pi)
	//fmt.Println("%.2f\n",math.Phi)

	//bool类型
	var isOK bool = false
	fmt.Println(isOK)

	//字符串类型
	var name string = "abc"
	fmt.Println(name)
	//字符类型
	var cake uint8 = 'a'
	var chineseName rune = '的'

	fmt.Println(cake, chineseName)

	//类型转换
	var age int8 = 12
	fmt.Println(float32(age))

	fmt.Println("int8 range: ", math.MinInt8, math.MaxInt8)
	var abv int32 = 1047483647
	fmt.Printf("int32: 0x%x  %d\n", abv, abv)
	//:=  注意 推导声明 左边的变量必须是没有定义过得
	// 如果是定义过的，那么他的类型就确定了，那么为什么还要编译器推导了，这不是浪费表情吗
	//出于这样的考虑，：= 左边的变量必须是没有定义过的
	bbba := int16(abv)
	fmt.Printf("int16: 0x%x  %d\n", bbba, bbba)

	//指针
	//指针，需要知道几个概念：指针地址、指针类型和指针取值
	//指针（pointer）概念在 Go 语言中被拆分为两个核心概念：
	//类型指针，允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。
	//切片，由指向起始元素的原始指针、元素数量和容量组成。

	girlName := "xiaoxiao"
	ptr := &girlName

	var pptr *string = ptr
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Println(pptr)

	//在对普通变量使用&操作符取地址获得这个变量的指针后，可以对指针使用*操作，也就是指针取值
	//ptr := &v    // v的类型为T
	//
	//其中 v 代表被取地址的变量，被取地址的 v 使用 ptr 变量进行接收，ptr 的类型就为*T，称做 T 的指针类型。*代表指针。
	fmt.Println("-----------------------------")
	people1 := 12
	people2 := 22
	Swap(&people1, &people2)
	fmt.Println(people1, people2)

}

func GetData() (int, int) {
	return 100, 2000
}

//指针实现变量交换
func Swap(a *int, b *int) {
	//把a指针所指向的值给一个临时变量temp
	temp := *a

	*a = *b

	*b = temp

}
