package main

import (
	"fmt"
	"math/rand"
	
	tls "saper/tools"
	rl "github.com/gen2brain/raylib-go/raylib"
)
// go build -ldflags="-H=windowsgui -s -w" -tags static -o Saper.exe
func main() {
	q,xrayfl,SLOZH,X,Y,SHAR := tls.MainMenu()
	if !q {return}
	
	score := 0
	sp := 22
	bg := rl.Gray
	loosefl := false
	winfl := false
	W := int32(10+sp*X )
	H := int32(25+sp*Y)

	btns := make([][]tls.Btn,Y) 
	//rasstanovka min + init polya
	for i := 0;i<Y;i++ {
		btns = append(btns, make([]tls.Btn,X))
		for j := 0;j<X;j++ {
			if rand.Intn(SLOZH) == 0 {
				btns[i] = append(btns[i], tls.NewBtn(5+j*sp,20+i*sp,"*"))
			} else {
				btns[i] = append(btns[i], tls.NewBtn(5+j*sp,20+i*sp,""))
			}
		}
	}
	// rasstanovka ciphor
	for i := 0;i<Y;i++ {
		for j := 0;j<X;j++ {
			if btns[i][j].Txt == "*" {continue}
			txt := tls.ParseBtn(j,i,&btns)
			if (txt == "0") {btns[i][j].Txt = ""} else {btns[i][j].Txt = txt}
		}
	}

	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(W,H,"Saper")
	rl.SetTargetFPS(30)
	rl.SetWindowIcon(*rl.LoadImage("icon.png"))

	rl.InitAudioDevice()
	scissors := rl.LoadSound("scissors.mp3")
	mine := rl.LoadSound("mine.mp3")
	win := rl.LoadSound("win.mp3")
	defer rl.UnloadSound(scissors)
	defer rl.UnloadSound(mine)
	defer rl.UnloadSound(win)

	sh := make([]tls.Sharik,SHAR)
	for i:=0;i<SHAR;i++ {
		rc := uint8(rand.Intn(255))
		sh[i] = tls.NewSharik(rand.Int31n(int32(W)),rand.Int31n(int32(H)),float32(rand.Intn(20)+10),rl.NewColor(rc,rc,rc,254),float32(rand.Intn(9)+1))
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(bg)

		msd := rl.GetMouseDelta()

		//shariki :p
		for i:=0;i<SHAR;i++{
			sh[i].Dr()
			sh[i].Mv(msd)
		}
		tls.UpdSh()
		//text
		rl.DrawText("v2.1",5,2,10,rl.Black)
		rl.DrawText("score: "+fmt.Sprint(score),W/2-40,2,20,rl.Black)
		rl.DrawText(fmt.Sprintf("%d:%d",X,Y),W-30,2,10,rl.Black)

		//otrisovka polya
		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
					btns[i][j].Dr(xrayfl)
			}
		}

		//obrabotka nazhatiy
		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
				if btns[i][j].IsCL() {
					if score == 0 && !btns[i][j].IsFlagged {if tls.ParseBtn(j,i,&btns) != "0" {btns[i][j].Txt = tls.ParseBtn(j,i,&btns)} else {btns[i][j].Txt = ""}}
					if btns[i][j].Txt == "*" && !btns[i][j].IsFlagged {loosefl = true}
					if btns[i][j].Txt == "" {tls.OpenZero(j,i,&btns);rl.PlaySound(scissors);continue}
					if !btns[i][j].IsOpen && !btns[i][j].IsFlagged {score++;btns[i][j].IsOpen = true;rl.PlaySound(scissors)}
					if tls.ParseWin(&btns) {winfl = true}
				}
				
				if btns[i][j].IsFl() {
					if !btns[i][j].IsFlagged {btns[i][j].IsFlagged = true;continue}
					if btns[i][j].IsFlagged {btns[i][j].IsFlagged = false}
				}
			}
		}

		// obrabotka smerty / pobedi
		if loosefl {
			rl.PlaySound(mine)
			bg = rl.Red
			rl.ClearBackground(bg)
			for i := 0;i<Y;i++ {
				for j := 0;j<X;j++ {
					btns[i][j].Dr(xrayfl)
				}
			}
			rl.EndDrawing()
			rl.WaitTime(2)
			rl.CloseAudioDevice()
			rl.CloseWindow()
			return
		} else if winfl {
			rl.PlaySound(win)
			bg = rl.Green
			rl.ClearBackground(bg)
			for i := 0;i<Y;i++ {
				for j := 0;j<X;j++ {
					btns[i][j].Dr(xrayfl)
				}
			}
			rl.EndDrawing()
			rl.WaitTime(3)
			rl.CloseAudioDevice()
			rl.CloseWindow()
			return
		}

		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
}