package main

import "golang.org/x/image/font"

type UserConfig struct {
	currentLanguage string
}

const (
	multiCharsPerSecond = 2
	gravity             = 1
)

const (
	screenWidth  = 1200
	screenHeight = 1024
)

const maxStartingHeight int = 20
const maxFallSpeed int = 6

const gameTitle = `gotype by sordidlist`

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)
