package main

import "image/color"

type MultiChar struct {
	inputChar            rune
	foreignChar          rune
	maskedChar           rune
	textColor            color.RGBA
	characterDisplayMode MultiCharDisplayMode
	xPos                 int
	yPos                 int
	moveSpeed            int
	timer                int
}

type MultiCharDisplayMode int8

const (
	ShowInputChar MultiCharDisplayMode = iota
	ShowForeignChar
	ShowMaskedChar
)
