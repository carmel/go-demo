package test

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := os.OpenFile("D:/tmp/excel/1.xlsx", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	xlsx, err := excelize.OpenReader(f)
	if err != nil {
		fmt.Println(err)
	}

	rows, _ := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			fmt.Println(data)
		}
	}
}
