package chapter2

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
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

	fmt.Println("---------------")

	//切片----很像python的切片操作
	//从数组切片
	var ints = [3]int{1, 2, 3}
	fmt.Println(ints, ints[:2])

	var highBuilding [30]int
	for i := 0; i < 30; i++ {
		highBuilding[i] = i + 1
	}

	//区间
	fmt.Println(highBuilding[10:15])
	//中间到尾部所有元素
	fmt.Println(highBuilding[20:])
	//开头到中间
	fmt.Println(highBuilding[:20])

	//表示原有切片
	var anotherInts = [4]int{2, 34, 5, 6}
	fmt.Println(anotherInts[:])
	fmt.Println(anotherInts[0:0])
	fmt.Println(anotherInts)

	//上面是从数组得到切片

	//直接声明新的切片
	var strList []string
	var numList []int
	var numListEmpty = []int{}

	fmt.Println(strList, numList, numListEmpty)
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	//make构造切片
	intsme := make([]int, 2, 10)
	intsme2 := make([]int, 2)
	fmt.Println(intsme, intsme2)
	fmt.Println(len(intsme), len(intsme2))

	//使用append() 为切片添加元素
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	var address []string
	address = append(address, "shenzheng")
	address = append(address, "beijing", "guangzhou", "nanchang")

	addresses := []string{"hangzhou", "shanghai"}
	address = append(address, addresses...)

	fmt.Println(address)

	//切片拷贝

	//相当于复制了一份，和以前的切面元素没关系了
	const elementCount = 1000
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData
	copyData := make([]int, elementCount)
	copy(copyData, refData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	fmt.Println(copyData[0])

	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	fmt.Println("-----------------")
	//切片删除元素---有点蠢啊，这都不提供内置方法
	seq := []string{"a", "b", "c", "d", "e"}
	//指定删除的位置
	removeIndex := 2
	fmt.Println(seq[:removeIndex], seq[removeIndex+1:])
	//将删除节点的前后 拼起来
	seq = append(seq[:removeIndex], seq[removeIndex+1:]...)
	fmt.Println("删除某元素的切片：", seq)

	fmt.Println("------------------")
	//map 映射，类似于hash表
	//Go 语言中 map 的定义是这样的：
	//map[KeyType]ValueType
	//
	//KeyType为键类型。
	//ValueType是键对应的值类型。
	//
	//一个 map 里，符合 KeyType 和 ValueType 的映射总是成对出现
	scene := make(map[string]int)
	scene["age"] = 12
	scene["name"] = 1234343
	fmt.Println(scene["name"])
	v, ok := scene["route"]
	fmt.Println(v)
	if !ok {
		fmt.Println("不存在")
	}

	//直接声明赋值map
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}

	//不限制值类型 interface{} 表示任意类型
	maps := make(map[string]interface{})
	maps["name"] = "xiaoqian"
	maps["age"] = 12
	maps["height"] = 167.3

	fmt.Printf("%+v", maps)
	fmt.Println()

	fmt.Println(m)

	//遍历map
	for k, v := range m {
		fmt.Println("key: ", k, "value: ", v)
	}

	//删除map元素 delete()
	delete(m, "S")
	fmt.Println(m)
	//释放map
	m = make(map[string]string)
	m = nil

	//并发下的map控制
	/*
		// 创建一个int到int的映射
			concurrentM := make(map[int]int)
			// 开启一段并发代码
			go func() {
				// 不停地对map进行写入
				for {
					concurrentM[1] = 1
				}
			}()
			// 开启一段并发代码
			go func() {
				// 不停地对map进行读取
				for {
					_ = concurrentM[1]
				}
			}()
			// 无限循环, 让并发程序在后台执行
			for {

			}
	*/

	//并发同步Map
	var syncMap sync.Map
	syncMap.Store("name", "xiaoqian")
	syncMap.Store("kk", 12)
	syncMap.Store("height", 180)

	fmt.Println(syncMap.Load("name"))
	syncMap.Delete("height")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("key: ", key, "value: ", value)
		return true
	})

	fmt.Println("--------------")

	//列表list（链表那种）
	var myList list.List
	myList.PushBack("a")
	myList.PushBack("b")
	fmt.Println(myList)

	myList2 := list.New()
	myList2.PushFront("abc")

	//s删除列表元素
	l := list.New()
	l.PushBack("cannon")
	l.PushFront("67")
	element := l.PushBack("first")
	fmt.Println(element)

	l.InsertBefore("the first?", element)
	fmt.Println(l)
	//遍历list中的元素
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	//遍历链表啊
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println("value: ", i.Value)
	}
}
