package demo

import (
	"fmt"
	"testing"
	"time"
)

//"bufio"

func TestFunc(t *testing.T) {
	str := "SUM(IF($C3=$C2,2,0)"
	al := []string{"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V"}
	for _, v := range al {
		str += fmt.Sprintf(",IF($%s3=$%s2,2,0)", v, v)
	}
	t.Error(str + ")")

	// 格式化
	const layout = "2006-01-02 15:04:05" //时间常量
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	cDate, _ := time.ParseInLocation(layout, "2020-05-12 00:00:00", Loc)

	for cDate.Format("2006-01-02") != "2020-06-12" {
		fmt.Println(cDate.Format("2006-01-02"))
		cDate = cDate.Add(86400000000000)
	}
}
