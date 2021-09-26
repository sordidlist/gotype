package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	target_language_keyboard_maps "gotype/target-language-keyboard-maps"
	"image/color"
	"log"
	"math/rand"
)

// Look for user configs, write if they don't exist
// Populate the progress bar
// If no user configs exist yet:
// Display banner and language selection menu
// User selects a language
// Display "Lessons" or "Challenges" menu

// If user selects "Lessons"
// Load any found lesson plugins
// Prompt the user with lesson options
// User selects a lesson
// Play the lesson

// If user selects "Challenges"
// Load any found challenge plugins
// Prompt the user with challenge options
// User selects a challenge
// Play the challenge

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

func init() {
	setFonts()
	seedRng()
}

func fetchNextMultiChar(g *Game) MultiChar {
	var kbMap keyboardMap
	if g.userConfig.currentLanguage == "Arabic" {
		kbMap = target_language_keyboard_maps.ArabicKeyboardMapType{}
	} else if g.userConfig.currentLanguage == "Russian" {
		kbMap = target_language_keyboard_maps.RussianKeyboardMapType{}
	}
	randomIndex := rand.Intn(kbMap.GetMapLength())
	count := 0
	var casedKbMap map[rune]rune
	upperOrLower := rand.Intn(1)
	if upperOrLower > 0 {
		casedKbMap = kbMap.GetUppercaseKeyboardMap()
	} else {
		casedKbMap = kbMap.GetLowercaseKeyboardMap()
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
				characterDisplayMode: ShowForeignChar,
				xPos:                 rand.Intn(screenWidth),
				yPos:                 50,
				moveSpeed:            4,
				timer:                0,
			}
			return randomMultiChar
		} else {
			count++
		}
	}
	return MultiChar{}
}

const (
	multiCharsPerSecond = 2
	gravity             = 2
)

func (g *Game) Update() error {

	// Populate a new character
	if g.counter%(ebiten.MaxTPS()/multiCharsPerSecond) == 0 {
		newMultiChar := fetchNextMultiChar(g)
		g.currentCharacters = append(g.currentCharacters, newMultiChar)
	}

	// Move all existing characters according to their properties
	if g.counter&ebiten.MaxTPS() == 0 {
		for charIndex := 0; charIndex < len(g.currentCharacters); charIndex++ {
			currentChar := g.currentCharacters[charIndex]
			currentChar.yPos += gravity * currentChar.moveSpeed
			currentChar.timer++
			g.currentCharacters[charIndex] = currentChar
		}
	}
	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for charIndex := 0; charIndex < len(g.currentCharacters); charIndex++ {
		currentChar := g.currentCharacters[charIndex]
		if currentChar.characterDisplayMode == ShowInputChar {
			text.Draw(screen, string(currentChar.inputChar), mplusNormalFont,
				currentChar.xPos, currentChar.yPos, currentChar.textColor)
		} else if currentChar.characterDisplayMode == ShowForeignChar {
			text.Draw(screen, string(currentChar.foreignChar), mplusNormalFont,
				currentChar.xPos, currentChar.yPos, currentChar.textColor)
		} else if currentChar.characterDisplayMode == ShowMaskedChar {
			text.Draw(screen, string(currentChar.maskedChar), mplusNormalFont,
				currentChar.xPos, currentChar.yPos, currentChar.textColor)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	config := UserConfig{
		currentLanguage: "Russian",
	}
	kbMap := getLanguageKeyboardMap(config.currentLanguage)
	if kbMap == nil {
		log.Fatal("Could not define keyboard map for language: " + config.currentLanguage)
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(gameTitle + " " + config.currentLanguage)
	if err := ebiten.RunGame(&Game{
		userConfig: config,
	}); err != nil {
		log.Fatal(err)
	}
}
