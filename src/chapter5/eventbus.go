package chapter5

import "fmt"

//实例恶化一个通过字符串映射的函数切片的map
var eventByName = make(map[string][]func(interface{}))

//注册时间提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	//找到存放函数的切片
	list := eventByName[name]
	//往切片中添加数据
	list = append(list, callback)
	//将新的切片保存回去
	eventByName[name] = list

}

//调用事件
func CallEvent(name string, params interface{}) {
	list := eventByName[name]

	for _, callback := range list {
		callback(params)
	}
}

//事件系统的使用
type Actor struct {
}

//为角色添加一个时间处理函数
func (a *Actor) OnEvent(params interface{}) {
	fmt.Println("actor event: ", params)
}

//全局时间
func GlobalEvent(params interface{}) {
	fmt.Println("global event: ", params)

}

func CallTheEventBusSystem() {
	actor := new(Actor)
	RegisterEvent("OnSkill", actor.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100)
}
