package main

import (
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

func main() {
	excelFileName := "C:/Users/Vector/Desktop/数据规范及接口说明.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		if len(sheet.Rows) != 0 {
			var buffer bytes.Buffer
			buffer.WriteString(`INSERT INTO `)
			buffer.WriteString(sheet.Name)
			buffer.WriteString(`(`)
			var title []string
			for _, cell := range sheet.Rows[0].Cells {
				title = append(title, cell.String())
			}
			buffer.WriteString(strings.Join(title, `,`))
			buffer.WriteString(`)VALUES(`)
			buffer.WriteString(strings.TrimSuffix(strings.Repeat(`?,`, len(title)), `,`))
			buffer.WriteString(`)`)

			fmt.Println(buffer.String())
			rows := sheet.Rows[1:]
			for _, row := range rows {
				fmt.Println()
				for _, cell := range row.Cells {
					text := cell.String()
					fmt.Printf("%s-", text)
				}
			}
		}
	}
}
