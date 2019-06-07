package chapter2

import (
	"fmt"
	"strconv"
)

func ContainerCakes() {
	//go容器测试

	//数组
	//声明
	var peoples [3]string
	peoples[0] = "xiaoqian"
	peoples[1] = "kangkang"
	peoples[2] = "cjk"

	//初始化
	//...表示让编译器确定数组大小。
	var cakes = [...]int{1, 2, 3}
	fmt.Println(cakes)
	fmt.Println("------------")

	//遍历数组
	var names = [3]string{"xiao", "zhu", "jing"}
	for k, v := range names {
		fmt.Println(k, v)
	}

	fmt.Println("--------------")
	for k, v := range cakes {
		fmt.Println(k, "---", v)
	}

	fmt.Println("--------------")
	for i := 0; i < len(peoples); i++ {
		peoples[i] = "people" + strconv.FormatInt(int64(i), 10)
	}

	fmt.Println(peoples)

}
