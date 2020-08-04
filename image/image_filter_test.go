package test

import (
	"fmt"
	"testing"

	"github.com/otiai10/gosseract"
	"gocv.io/x/gocv"
)

func TestMain(t *testing.T) {
	// imgMatrix := imgo.MustRead("1.jpg")
	// //获取一个[][][]uint8对象

	// //binaryzation process of image matrix , threshold can use 127 to test
	// //func Binaryzation(src [][][]uint8, threshold int) [][][]uint8{}
	// imgMatrixGray := imgo.Binaryzation(imgMatrix, 127)

	// grayRgb := "gray_rgb.png"
	// err := imgo.SaveAsPNG(grayRgb, imgMatrixGray)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%s generate finish\n", grayRgb)
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("chi_sim")
	client.SetImage("1.jpg")
	text, _ := client.Text()
	fmt.Println(text)
}

func TestCV(t *testing.T) {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
