package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattn/go-tty"
	target_language_keyboard_maps "gotype/target-language-keyboard-maps"
	"image/color"
	"log"
	"math/rand"
)

func populateMultiChars(g *Game) {
	if g.counter%(ebiten.MaxTPS()/multiCharsPerSecond) == 0 {
		newMultiChar, err := fetchNextMultiChar(g)
		if err != nil {

		} else {
			g.currentCharacters = append(g.currentCharacters, newMultiChar)
		}
	}
	g.counter++
}

func affectAllMultiChars(g *Game) {
	for charIndex := 0; charIndex < len(g.currentCharacters); charIndex++ {
		currentChar := g.currentCharacters[charIndex]
		currentChar.yPos += gravity * currentChar.moveSpeed
		currentChar.timer++
		g.currentCharacters[charIndex] = currentChar
	}
}

func updateMultiCharsFromUserInput(g *Game) {
	inputRunes := ebiten.InputChars()
	if len(inputRunes) > 0 {
		for inputRuneIndex := 0; inputRuneIndex < len(inputRunes); inputRuneIndex++ {
			for charIndex := 0; charIndex < len(g.currentCharacters); charIndex++ {
				currentChar := g.currentCharacters[charIndex]
				if currentChar.inputChar == inputRunes[inputRuneIndex] {
					if currentChar.characterDisplayMode == ShowInputChar {
						currentChar.characterDisplayMode = ShowForeignChar
						g.currentCharacters[charIndex] = currentChar
						break
					}
				}
			}
		}
	}
}

type keyboardMapCase int8

const (
	UppercaseKeyboardMap keyboardMapCase = iota
	LowercaseKeyboardMap
	NumbersKeyboardMap
)

func fetchNextMultiChar(g *Game) (MultiChar, error) {
	var kbMap keyboardMap
	if g.userConfig.currentLanguage == "Arabic" {
		kbMap = target_language_keyboard_maps.ArabicKeyboardMapType{}
	} else if g.userConfig.currentLanguage == "Russian" {
		kbMap = target_language_keyboard_maps.RussianKeyboardMapType{}
	} else {
		return MultiChar{}, errors.New("could not load a keyboard map for language: " + g.userConfig.currentLanguage)
	}
	randomIndex := rand.Intn(kbMap.GetMapLength())
	count := 0
	var casedKbMap map[rune]rune
	casedKbMode := keyboardMapCase(rand.Intn(3))

	if casedKbMode == UppercaseKeyboardMap {
		casedKbMap = kbMap.GetUppercaseKeyboardMap()
	} else if casedKbMode == LowercaseKeyboardMap {
		casedKbMap = kbMap.GetLowercaseKeyboardMap()
	} else if casedKbMode == NumbersKeyboardMap {
		casedKbMap = kbMap.GetShiftNumberRowMap()
	}
	for key, value := range casedKbMap {
		if count == randomIndex {
			randomMultiChar := MultiChar{
				inputChar:   key,
				foreignChar: value,
				maskedChar:  '*',
				textColor: color.RGBA{
					R: uint8(rand.Intn(255)),
					G: uint8(rand.Intn(255)),
					B: uint8(rand.Intn(255)),
					A: uint8(rand.Intn(255)),
				},
				characterDisplayMode: ShowInputChar,
				xPos:                 rand.Intn(screenWidth),
				yPos:                 50,
				moveSpeed:            2,
				timer:                0,
			}
			return randomMultiChar, nil
		} else {
			count++
		}
	}
	return MultiChar{}, errors.New("failed to return a new multi char")
}

func readRunes(tty *tty.TTY) {
	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		printTextColoredText256(string(r), rand.Intn(255))
	}
}
