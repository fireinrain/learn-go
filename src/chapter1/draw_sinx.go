package chapter1

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func RunDrawSinx() {
	//图片大小
	const size = 300

	//创建一个灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	//像素填充
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2

		pic.SetGray(x, int(y), color.Gray{0})
	}

	//创建文件
	file, error := os.Create("sin.png")
	if error != nil {
		log.Fatal(error)
	}
	png.Encode(file, pic)
	file.Close()

}
