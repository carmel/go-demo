package main

type Painter interface {
	Draw()
}
type Cowboy interface {
	Draw()
}

type Xiaowang struct {
}
type XiaowangAsPainter Xiaowang

func (p *XiaowangAsPainter) Draw() {
	//画画
}

type XiaowangAsCowboy Xiaowang

func (p *XiaowangAsCowboy) Draw() {
	// 拔枪
}
func main() {
	var xw Xiaowang
	var painter Painter = (*XiaowangAsPainter)(&xw)
	painter.Draw()
	var cowboy Cowboy = (*XiaowangAsCowboy)(&xw)
	cowboy.Draw()
}
