// Package main shows imgui demo window.
package main

import (
	"github.com/AllenDang/cimgui-go/imgui"

	g "github.com/AllenDang/giu"
)

func loop() {
	imgui.ShowDemoWindow()
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0)
	wnd.Run(loop)
}
