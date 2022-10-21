package demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/carmel/xlsx"
)

func TestXlsx(t *testing.T) {
	f, err := os.OpenFile("D:/tmp/excel/1.xlsx", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	xlsx, err := xlsx.OpenReader(f)
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
