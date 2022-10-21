package robotgo

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestRobotHandler(t *testing.T) {
	fpid, err := robotgo.FindIds("cmd")
	if err == nil {
		fmt.Println("pids...", fpid)

		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])

			robotgo.Kill(fpid[0])
		}
	}

	robotgo.ActiveName("chrome")

	isExist, err := robotgo.PidExists(100)
	if err == nil && isExist {
		fmt.Println("pid exists is", isExist)

		robotgo.Kill(100)
	}

	abool := robotgo.ShowAlert("test", "robotgo")
	if abool {
		fmt.Println("ok@@@ ", "ok")
	}

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
}

func TestKeyboard(t *testing.T) {
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("だんしゃり", 1.0)
	// robotgo.TypeString("テストする")

	robotgo.TypeStr("Hi galaxy. こんにちは世界.")
	robotgo.Sleep(1)

	// ustr := uint32(robotgo.CharCodeAt("Test", 0))
	// robotgo.UnicodeType(ustr)

	robotgo.KeyTap("enter")
	// robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")

	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("Test")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
}

func TestMouse(t *testing.T) {
	// fy := 205
	fy := 253

	for i := 0; i < 9; i++ {
		robotgo.Move(1588, fy+i*96) // 相隔96
		robotgo.Click()
		robotgo.Sleep(1)

		robotgo.Move(1095, 317)
		robotgo.Click() // 导出
		robotgo.Sleep(1)

		robotgo.KeyTap("enter") // 触发enter键之前需要休息1秒钟；保存文件
		robotgo.Sleep(1)

		robotgo.Move(1107, 592) // 确定导出
		robotgo.Sleep(1)
		robotgo.Click()

		robotgo.Move(1133, 226) // 关闭
		robotgo.Sleep(1)
		robotgo.Click()
		// robotgo.Sleep(1)
	}

	// robotgo.MoveRelative(10, -200)

	// // move the mouse to 100, 200
	// robotgo.MoveMouse(100, 200)

	// robotgo.Drag(10, 10)
	// robotgo.Drag(20, 20, "right")
	// //
	// robotgo.DragSmooth(10, 10)
	// robotgo.DragSmooth(100, 200, 1.0, 100.0)

	// // smooth move the mouse to 100, 200
	// robotgo.MoveSmooth(100, 200)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	// robotgo.MoveSmoothRelative(10, -100, 1.0, 30.0)

	// for i := 0; i < 1080; i += 1000 {
	// 	fmt.Println(i)
	// 	robotgo.MoveMouse(800, i)
	// }
}

func TestBitmap(t *testing.T) {
	bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)

	robotgo.SaveBitmap(bitmap, "test.png")
}
