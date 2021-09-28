package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func loadOpentypeFont(fontPath string) (font.Face, error) {
	fontFile, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	otFont, err := opentype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	faceOptions := &opentype.FaceOptions{
		Size:    50,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}
	otFontFace, err := opentype.NewFace(otFont, faceOptions)
	if err != nil {
		return nil, err
	}
	return otFontFace, nil
}

// All fonts downloaded from https://fonts.google.com
func setFonts(config *UserConfig) {
	//_, err := loadOpentypeFont("assets/ScheherazadeNew-Regular.ttf")
	scheherazadeFont, err := loadOpentypeFont("fonts/ScheherazadeNew-Bold.ttf")
	//_, err := loadOpentypeFont("assets/Amiri-Regular.ttf")
	amiriFont, err := loadOpentypeFont("fonts/Amiri-Bold.ttf")
	//_, err := loadOpentypeFont("assets/Roboto-Regular.ttf")
	robotoFont, err := loadOpentypeFont("fonts/Roboto-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}

	inputCharacterFont = scheherazadeFont
	if config.currentLanguage == ARABIC_LANGUAGE {
		foreignCharacterFont = amiriFont
	} else if config.currentLanguage == RUSSIAN_LANGUAGE {
		foreignCharacterFont = robotoFont
	}
}

func seedRng() {
	rand.Seed(time.Now().UnixNano())
}
