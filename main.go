package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func onMenuButton() {
	fmt.Println("menu bar click")
}

func loop() {
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.Button("Button inside menu").OnClick(onMenuButton),
			),
		),
		g.Label("Hello world from giu"),
		g.Row(
			g.Button("Click Me").OnClick(onClickMe),
			g.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}