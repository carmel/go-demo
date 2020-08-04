package svg_test

import (
	"go-demo/svg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

const (
	DECORATE_SPACE = 4 // 音符的修饰符与之间距
	NOTE_SPACE     = 9 // 每个音符之间的间距
	LINE_SPACE     = 16
)

func TestSvg(t *testing.T) {

	http.Handle("/svg", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		canvas := svg.New(w)
		canvas.StartPercentView(100, 0, 0, 0, 210, 297, "xMidYMid meet")
		canvas.Text(105, 18, "义勇军进行曲", "font-size:0.645rem;font-weight:bold;text-anchor:middle")
		canvas.Text(10, 28, "G 2/4", "font-size:0.445rem;font-weight:bold;text-anchor:start")
		canvas.Text(200, 28, "田汉 词", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Text(200, 38, "聂耳 曲", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Text(10, 38+LINE_SPACE, "1", "font-size:0.445rem;font-weight:bold;text-anchor:start")
		canvas.Text(10+NOTE_SPACE, 38+LINE_SPACE, "•", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Text(10+2*NOTE_SPACE, 38+LINE_SPACE, "1", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Text(10+3*NOTE_SPACE, 38+LINE_SPACE, "1", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Text(10, 38+2*LINE_SPACE, "1", "font-size:0.445rem;font-weight:bold;text-anchor:end")
		canvas.Arc(105, 105, 10, 4, 0, false, true, 125, 105, `id="top"`, `fill="none"`, `stroke="red"`)
		canvas.Translate(300, 100)
		canvas.Line(-7.478125, -16.15, 7.478125, -16.15)
		canvas.Gend()
		// <g transform="matrix(1,0,0,1,56.74,0)"><use xlink:href="#p502"></use><line x1="0" x2="51.45899375" y1="0" y2="0" style="stroke-width: 1;"></line><line x1="0" x2="51.45899375" y1="-3.4" y2="-3.4" style="stroke-width: 1;"></line><path d="M4.939724922180176,-19.947500228881836c0,-8 97.0532390625,-8 97.0532390625,0c0,-10 -97.0532390625,-10 -97.0532390625,0"></path></g>

		canvas.End()
	}))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}

func TestAm(t *testing.T) {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./"))))
	http.Handle("/svg", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		canvas := svg.New(w)
		var width, height float32 = 600.0, 600.0
		var rsize float32 = 20.0
		csize := rsize / 2
		var duration float32 = 5.0
		var repeat float32 = 2
		canvas.Start(width, height)
		canvas.Image(0, 0, 100, 140, "res/gopher.jpg", `id="gopher"`)
		canvas.Arc(0, 250, 10, 10, 0, false, true, 500, 250, `id="top"`, `fill="none"`, `stroke="red"`)
		canvas.Arc(0, 250, 10, 10, 0, true, false, 500, 250, `id="bot"`, `fill="none"`, `stroke="blue"`)
		canvas.Circle(0, 0, csize, `fill="red"`, `id="red-dot"`)
		canvas.Circle(0, 0, csize, `fill="blue"`, `id="blue-dot"`)
		canvas.AnimateMotion("#red-dot", "#top", 5, repeat)
		canvas.AnimateMotion("#blue-dot", "#bot", 3, repeat)
		canvas.AnimateMotion("#gopher", "#top", duration, repeat)
		canvas.End()
	}))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}

func TestAmt(t *testing.T) {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./"))))
	http.Handle("/svg", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		canvas := svg.New(w)
		var width, height float32 = 500.0, 500.0
		var duration, repeat float32 = 1.0, 0

		canvas.Start(width, height)
		canvas.Rect(0, 0, width, height)

		// Translate
		canvas.CenterRect(0, 0, 40, 30, "fill:red", `id="redbox"`)
		canvas.CenterRect(0, 0, 40, 30, "fill:blue", `id="bluebox"`)
		canvas.AnimateTranslate("#redbox", 100, 100, 200, 200, duration, repeat)
		canvas.AnimateTranslate("#bluebox", 200, 200, 100, 100, duration, repeat)

		// Scale
		canvas.Translate(200, 100)
		canvas.CenterRect(0, 0, 40, 30, "fill:green", `id="greenbox"`)
		canvas.Gend()
		canvas.AnimateScale("#greenbox", 1, 3, duration, repeat)

		// SkewX
		canvas.Translate(300, 100)
		canvas.CenterRect(0, 0, 40, 30, "fill:purple", `id="purplebox"`)
		canvas.Gend()
		canvas.AnimateSkewX("#purplebox", 0, 45, duration, repeat)

		// SkewY
		canvas.Translate(400, 100)
		canvas.CenterRect(0, 0, 40, 30, "fill:lightsteelblue", `id="lsbox"`)
		canvas.Gend()
		canvas.AnimateSkewY("#lsbox", 0, 45, duration, repeat)

		// Rotate
		canvas.Translate(width/4, height/2)
		canvas.CenterRect(0, 0, 40, 30, "fill:maroon", `id="rotbox"`)
		canvas.Gend()
		canvas.AnimateRotate("#rotbox", 0, 75, 75, 360, 75, 75, duration*2, repeat)

		canvas.End()
	}))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}

func TestAnimate(t *testing.T) {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./"))))
	http.Handle("/svg", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		canvas := svg.New(w)
		var width, height int = 500, 500
		var rsize float32 = 100
		csize := rsize / 2
		var duration float32 = 5.0
		var repeat float32 = 5
		var imw, imh float32 = 100, 144
		canvas.Start(float32(width), float32(height))
		canvas.Circle(csize, csize, csize, `fill="red"`, `id="circle"`)
		canvas.Image((float32(width)/2)-(imw/2), 0, imw, imh, "res/gopher.jpg", `id="gopher"`)
		canvas.Square(float32(width)-rsize, 0, rsize, `fill="blue"`, `id="square"`)
		canvas.Animate("#circle", "cx", 0, width, duration, repeat)
		canvas.Animate("#circle", "cy", 0, height, duration, repeat)
		canvas.Animate("#square", "x", width, 0, duration, repeat)
		canvas.Animate("#square", "y", height, 0, duration, repeat)
		canvas.Animate("#gopher", "y", 0, height, duration, repeat)
		canvas.End()
	}))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}
