package chapter1

import "fmt"

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

}

func GetData() (int, int) {
	return 100, 2000
}
