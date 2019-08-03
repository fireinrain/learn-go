package chapter6

import (
	"encoding/json"
	"fmt"
)

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Wheel
	Engine
}

type A struct {
	a int
}
type B struct {
	a int
}
type C struct {
	A
	B
}

//定义手机屏幕
type Screen struct {
	Size       float32
	ResX, ResY int
}

//定义电池
type Battery struct {
	Capacity int
}

func genJsonDate() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchId bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery:    Battery{Capacity: 2910},
		HasTouchId: false,
	}

	marshal, _ := json.Marshal(raw)
	return marshal
}

func GetACar() {

	car := Car{
		Wheel: Wheel{
			Size: 12,
		},
		Engine: Engine{
			Type:  "Monster",
			Power: 100,
		},
	}

	c := C{
		A: A{a: 1},
		B: B{a: 2},
	}
	c.A.a = 13
	fmt.Println(c)

	fmt.Printf("%+v\n", car)

	//序列 反序列化JSON 数据
	data := genJsonDate()
	fmt.Println(string(data))

	//反序列化
	sc := struct {
		Screen
		HasTouchId bool
	}{}
	json.Unmarshal(data, &sc)

	fmt.Printf("%+v\n", sc)

}
