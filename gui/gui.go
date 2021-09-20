package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

func CreateNewGUI() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
}
