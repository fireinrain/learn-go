package chapter3

import "fmt"

func ControlTheProcess() {
	//流程控制
	var ten int = 10
	if ten > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}

	//特殊写法
	//if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：
	//if err := Connect(); err != nil {
	//    fmt.Println(err)
	//    return
	//}
	//Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
	//
	//err！=nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。

	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	for ; ; step++ {
		if step > 5 {
			break
		}
		fmt.Println("step: ", step)
	}

	fmt.Println(step)
	for {
		if step > 10 {
			break
		}

		step++
		fmt.Println(step)
	}

	for step < 100 {
		step++
		if step == 99 {
			fmt.Println(step)
		}
	}
	fmt.Println("-------------")

	start := 0
	end := 0
	//输出九九乘法表
	for start < 9 {
		start++
		for end < start {
			end++
			fmt.Print(start, "*", end, "=", start*end, " ")
		}
		end = 0
		fmt.Println()
	}

	//输出九九乘法表2
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}

	//键值对 遍历
	//Go 语言可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有一定的规律：
	//数组、切片、字符串返回索引和值。
	//map 返回键和值。
	//通道（channel）只返回通道内的值
	for key, value := range []int{1, 2, 4} {
		fmt.Println(key, value)
	}

	//遍历字符串
	var strs = "hello 你好"
	for key, value := range strs {
		fmt.Printf("key:%d,value: 0x%x\n", key, value)
	}

	//遍历map
	m := make(map[string]int)
	m["hello"] = 100
	m["world"] = 200

	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}

	//遍历管道
	channels := make(chan int)

	go func() {
		channels <- 1
		channels <- 2
		channels <- 3
		//如果不关闭的话会报错
		//goroutine 执行完了，管道不会有值
		//但是这是后主gorutine 还期望有值，所以会报错，因为等不到
		close(channels)
	}()

	for v := range channels {
		fmt.Println("channel push: ", v)
	}

	//丢弃某些遍历值
	maps := map[string]int{
		"hellos":   123,
		"catch me": 345,
	}
	for _, value := range maps {
		fmt.Println(value)
	}

	for key, _ := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d \n", key)
	}
	fmt.Println("--------------")

	//switch 不用想java 那样break
	var hello = "hello"
	switch hello {
	case "hello":
		fmt.Println("world")

	case "world":
		fmt.Println("haha")

	default:
		fmt.Println("i am a  default")
	}

	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {

	case r > 10 && r < 20:
		fmt.Println(r)
	}

	//goto 什么鬼 啊
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")

	//break

outerloop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break outerloop

			case 3:
				fmt.Println(i, j)
				break outerloop
			}
		}
	}

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
