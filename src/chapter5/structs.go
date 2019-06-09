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
