package tls

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Btn struct {
	X int32;
	Y int32;
	Txt string;
	Sz int32 `default: 20`; //TODO a to poka ni na chto ne vliaet
	IsOpen bool `default: false`;
	IsFlagged bool `default: false`;
}

var BtnSz int32 = 20
var cntr float64 = 0

func NewBtn(x int,y int,txt string) Btn {
	return Btn{X:int32(x),Y:int32(y),Txt:txt}
}

func (b Btn) Pr() {
	fmt.Println(b.X,b.Y,b.Txt,b.IsOpen)
}

// otrisovka polya
func (b Btn) Dr(xrayfl bool) {
	if b.IsOpen {
		rl.DrawRectangle(b.X-1,b.Y-1,BtnSz+1,BtnSz+1,rl.Black)
		rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.LightGray)
		rl.DrawText(b.Txt,b.X+5,b.Y,22,rl.Black)
	} else if b.IsFlagged {
		rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.Red)
	} else {
		rl.DrawRectangle(b.X-1,b.Y-1,BtnSz+1,BtnSz+1,rl.Black)
		rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.DarkGray)
	}
	if !b.IsOpen && xrayfl {rl.DrawText(b.Txt,b.X+5,b.Y,22,rl.Black)}
}

// func (b Btn) DrBtn() {
// 	b.dr(false)
// }

// func (b Btn) DrBtnXR() {
// 	b.dr(true)
// }

// skolko min sosedei
func ParseBtn(x int,y int,arr *[][]Btn) string {
	cntr := 0
	for i := y-1;i<=y+1;i++{
		for j := x-1;j<=x+1;j++{
			if i < 0 || i > len((*arr))-1 || j < 0 || j > len((*arr)[i])-1 {continue}
			if (*arr)[i][j].Txt == "*" {cntr++}
		}
	}
	return fmt.Sprintf("%d",cntr)
}

// esli nazhali na pole levoi
func (b Btn) IsCL() bool {
	ms := rl.GetMousePosition()
	return rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ms.X	> float32(b.X) && ms.X	< float32(b.X+BtnSz) && ms.Y > float32(b.Y) && ms.Y < float32(b.Y+BtnSz)
}

// esli nazhali na pole pravoi
func (b Btn) IsFl() bool {
	ms := rl.GetMousePosition()
	return rl.IsMouseButtonPressed(rl.MouseButtonRight) && ms.X	> float32(b.X) && ms.X	< float32(b.X+BtnSz) && ms.Y > float32(b.Y) && ms.Y < float32(b.Y+BtnSz)
}

// otkrit sosedei
func openAllNear(x int,y int,arr *[][]Btn) {
	for i := y-1;i<=y+1;i++{
		for j := x-1;j<=x+1;j++{
			if i < 0 || i > len((*arr))-1 || j < 0 || j > len((*arr)[i])-1 || (*arr)[i][j].Txt == "*" {continue}
			(*arr)[i][j].IsOpen = true
		}
	}
}

// avto otkritie pustih sosedei cherez dfs
func OpenZero(x int,y int,arr *[][]Btn) {
	if y < 0 || y > len((*arr))-1 || x < 0 || x > len((*arr)[0])-1 || (*arr)[y][x].Txt != "" || (*arr)[y][x].IsOpen == true {return}
	(*arr)[y][x].IsOpen = true

	OpenZero(x+1,y,arr)
	OpenZero(x-1,y,arr)
	OpenZero(x,y-1,arr)
	OpenZero(x,y+1,arr)
	OpenZero(x+1,y+1,arr)
	OpenZero(x-1,y-1,arr)
	OpenZero(x+1,y-1,arr)
	OpenZero(x-1,y+1,arr)

	openAllNear(x,y,arr)
}

func ParseWin(arr *[][]Btn) bool {
	for _,q := range *arr{
		for _,w := range q{
			if w.Txt != "*" && !w.IsOpen {return false}
		}
	}
	return true
}

//tipa sharik na bg
type Sharik struct {
	X int32;
	Y int32;
	R float32;
	C rl.Color;
	Weight float32;
}

func NewSharik(x int32, y int32, r float32, c rl.Color,w float32) Sharik {
	return Sharik{X:x,Y:y,R:r,C:c,Weight:w}
}

func (s Sharik) Dr() {
	rl.DrawCircle(s.X,s.Y,s.R+float32(math.Abs(math.Sin(cntr)*float64(s.R))),s.C)
}

func (s *Sharik) Mv(v rl.Vector2) {
	s.X -= int32(v.X/s.Weight)
	s.Y -= int32(v.Y/s.Weight)
	q := (v.X+v.Y)/10
	// stranno vedut sebia s etim codom
	// if q > 254 {q = 254}
	// if q < 0 {q = 0}
	s.C.R -= uint8(q)
	s.C.G -= uint8(q)
	s.C.B -= uint8(q)
}

func UpdSh() {
	cntr += 0.01
	if math.Abs(math.Sin(cntr)-math.Sin(0)) < 0.05 {cntr = 0} // esli primerno ravno 0 to obnuliaem ot stackoverflow
}

// otrisovka checkboxa
func checkBox(hb rl.Rectangle, text string,txtsz int32,checked bool) bool {
	ms := rl.GetMousePosition()
	b := hb.ToInt32()
	rl.DrawText(text, b.X+b.Width+10, b.Y+b.Height/5, txtsz, rl.Black)
	ispr := rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ms.X	> float32(b.X) && ms.X	< float32(b.X+hb.ToInt32().Width) && ms.Y > float32(b.Y) && ms.Y < float32(b.Y+hb.ToInt32().Height)
	if !checked {rl.DrawRectangleLinesEx(hb,2,rl.Black)} else {rl.DrawRectangleRec(hb,rl.Black)} 
	
	return ispr
}

// menu vibora
func MainMenu() (bool,bool,int,int,int,int) { // submit/net + xray + slozh + x + y + shariki
	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(200,400,"Saper")
	rl.SetTargetFPS(30)
	rl.SetWindowIcon(*rl.LoadImage("icon.png"))
	var xrch,mch,hch,hdch bool
	ech := true
	recs := []rl.Rectangle{rl.NewRectangle(10,10,30,30),rl.NewRectangle(10,50,30,30),rl.NewRectangle(10,90,30,30),rl.NewRectangle(10,130,30,30),rl.NewRectangle(10,170,30,30)}
	subr := rl.NewRectangle(10,210,38,20)

	sh := make([]Sharik,10)
	for i:=0;i<10;i++ {
		rc := uint8(rand.Intn(255))
		sh[i] = NewSharik(rand.Int31n(200),rand.Int31n(400),float32(rand.Intn(20)+10),rl.NewColor(rc,rc,rc,254),float32(rand.Intn(9)+1))
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		ms := rl.GetMousePosition()
		msd := rl.GetMouseDelta()
		
		rl.ClearBackground(rl.Gray)
		//shariki :p
		for i:=0;i<10;i++{
			sh[i].Dr()
			sh[i].Mv(msd)
		}
		UpdSh()
		
		//checkboxi
		if checkBox(recs[0],"Xray-vision",20,xrch) {xrch = !xrch}
		if checkBox(recs[1],"Easy",20,ech) {ech = true;mch = false;hdch = false;hch = false;}
		if checkBox(recs[2],"Medium",20,mch) {ech = false;mch = true;hdch = false;hch = false;}
		if checkBox(recs[3],"Hard",20,hch) {ech = false;mch = false;hdch = false;hch = true;}
		if checkBox(recs[4],"Harder",20,hdch) {ech = false;mch = false;hdch = true;hch = false;}
		
		//knopka nachala
		rl.DrawRectangleRec(subr,rl.Black)
		rl.DrawRectangleLinesEx(subr,2,rl.White)
		rl.DrawText("start",15,215,10,rl.White)
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ms.X	> float32(subr.X) && ms.X < float32(subr.X+subr.Width) && ms.Y > float32(subr.Y) && ms.Y < float32(subr.Y+subr.Height) {
			rl.CloseWindow()
			if ech{return true,xrch,10,10,10,7+rand.Intn(3)-3}
			if mch {return true,xrch,7,20,20,10+rand.Intn(3)-3}
			if hch {return true,xrch,5,25,25,15+rand.Intn(3)-3}
			if hdch{return true,xrch,3,30,30,20+rand.Intn(3)-3}
		}
		
		rl.EndDrawing()
	}
	return false,false,7,20,20,7
}