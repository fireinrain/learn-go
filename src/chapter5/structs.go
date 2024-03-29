package chapter5

import "fmt"

func StructsPlay() {

	//结构体
	//结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存

	//创建结构体实例的方法有多种
	//1 基本方法 结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化。
	var p Point
	p.X = 132
	p.Y = 1
	fmt.Println(p.X, p.Y)

	//2.  创建指针类型的结构体
	player := new(Player)
	player.Name = "cake"
	player.HealthPoint = 12

	fmt.Printf("palyers: %+v", player)
	fmt.Println()
	//3 取结构体的地址实例化
	// 取地址实例化是最广泛的一种结构体实例化方式。可以使用函数封装上面的初始化过程
	var versionInt int = 1
	command := &Command{}
	command.Name = "start"
	command.Var = &versionInt
	command.Comment = "开始启动"

	fmt.Println(command)

	fmt.Println("----------")
	people := &People{
		name: "xiqoain",
		child: &People{
			name: "baba",
			child: &People{
				name: "wo",
			},
		},
	}
	fmt.Printf("people: %+v", people)

	fmt.Println()
	fmt.Println("------")
	i := &People{
		name:  "xiaoxiao",
		child: new(People),
	}

	fmt.Println(i)

	//初始化匿名结构体
	ins := &struct {
		name string
		age  int32
	}{
		name: "lily",
		age:  123,
	}

	fmt.Println(ins)

	// 实例化一个匿名结构体
	msg := &struct {
		// 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printNoNameStruct(msg)

	var cake string = "xiaoxi"
	var cakePtr *string = &cake
	fmt.Println(cakePtr)
	fmt.Println("---------")

	//go中的构造函数
	//Go 语言的类型或结构体没有构造函数的功能。结构体的初始化过程可以使用函数封装实现。
	//
	//其他编程语言构造函数的一些常见功能及特性如下：
	//每个类可以添加构造函数，多个构造函数使用函数重载实现。
	//构造函数一般与类名同名，且没有返回值。
	//构造函数有一个静态构造函数，一般用这个特性来调用父类的构造函数。
	//对于 C++ 来说，还有默认构造函数、拷贝构造函数等。

	//模拟构造函数重载

	//模拟父级构造函数
	cat := NewBlackCat("yellow")
	fmt.Println(cat.Color)

	//方法 与接收器
	bag := &Bag{}
	fmt.Println(*bag)
	fmt.Println(bag)
	Insert(*bag, 2)

	fmt.Printf("----------")
	fmt.Println()

	//指针类型的接收器
	property := &Property{}
	property.SetValue(1)
	fmt.Println(property.value)
	//在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器。
	// 大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。

	//Go 语言可以对任何类型添加方法。给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型。

	var myint myInt = 12
	fmt.Println(myint.Is0())

	fmt.Println("---------")
	//方法与函数的统一调用

	//声明一个函数回调
	var delegate func(int)
	//创建结构体实例
	c := &class{}
	//将回调设置为c的Do方法
	delegate = c.Do
	//调用
	delegate(100)
	//将回调设置为普通函数
	delegate = Do
	//调用
	delegate(120)
	//无论是普通函数还是结构体的方法，
	// 只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法

}

//定义结构体
type class struct {
}

//给结构体添加DO方法
func (c *class) Do(v int) {
	fmt.Println("call method do: ", v)
}

//普通函数
func Do(v int) {
	fmt.Println("call func do: ", v)
}

type myInt int

func (i myInt) Is0() bool {
	return i == 0
}
func (p *Property) Value() int {
	return p.value
}
func (p *Property) SetValue(value int) {
	p.value = value
}

type Property struct {
	value int
}

func Insert(bag Bag, item int) {
	ints := append(bag.item, item)
	fmt.Println(ints)
}

type Bag struct {
	item []int
}

//模拟父级构造函数
type BlackCat struct {
	Cat
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

//模拟构造函数重载
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: "red",
	}
}

type Cat struct {
	Color string
	Name  string
}

// // 使用动词%T打印noname的类型
func printNoNameStruct(nonamestruct *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", nonamestruct)
}

type People struct {
	name  string
	child *People
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Point struct {
	X int
	Y int
}

type Color struct {
	R, G, B byte
}
