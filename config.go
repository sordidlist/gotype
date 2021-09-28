package main

import "golang.org/x/image/font"

type UserConfig struct {
	currentLanguage string
}

const (
	multiCharsPerSecond = 1
	gravity             = 1
)

const (
	screenWidth  = 1200
	screenHeight = 1024
	bufferWidth  = 30
)
const dpi = 72

const maxStartingHeight int = 20
const maxFallSpeed int = 6

const ARABIC_LANGUAGE = "Arabic"
const RUSSIAN_LANGUAGE = "Russian"

const multiCharMaxLifetime = 500

const gameTitle = `gotype by sordidlist`

var (
	inputCharacterFont   font.Face
	foreignCharacterFont font.Face
	amiriFont            font.Face
	robotoFont           font.Face
	scheherazadeFont     font.Face
)
