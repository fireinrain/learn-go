package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	PlayWithExcel()

}

//测试excel
func PlayWithExcel() {
	readExcel()
	//writeExcel()
}

func writeExcel() {
	file := excelize.NewFile()

	index := file.NewSheet("myGosheet")
	file.SetCellValue("myGosheet", "A1", "Hello Go")
	file.SetCellValue("Sheet1", "B2", 100)

	file.SetActiveSheet(index)
	e := file.SaveAs("./go.xlsx")
	if e != nil {
		fmt.Println(e)
	}

}

func readExcel() {
	file, e := excelize.OpenFile("./go.xlsx")
	if e != nil {
		fmt.Println(e)
		return
	}
	cellValue, err := file.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("B2: ", cellValue)

	rows, err := file.GetRows("myGosheet")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
