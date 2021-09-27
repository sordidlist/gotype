package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/mattn/go-tty"
	"log"
	"os"
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

func init() {
	setFonts()
	seedRng()
}

func (g *Game) Update() error {
	populateMultiChars(g)
	updateMultiCharsFromUserInput(g)
	affectAllMultiChars(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for charIndex := 0; charIndex < len(g.currentCharacters); charIndex++ {
		currentChar := g.currentCharacters[charIndex]
		if currentChar.characterDisplayMode == ShowInputChar {
			text.Draw(screen, string(currentChar.inputChar), mplusBigFont,
				currentChar.xPos, currentChar.yPos, currentChar.textColor)
		} else if currentChar.characterDisplayMode == ShowForeignChar {
			text.Draw(screen, string(currentChar.foreignChar), mplusBigFont,
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
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	go readRunes(tty)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(gameTitle + " " + config.currentLanguage)
	if err := ebiten.RunGame(&Game{
		userConfig: config,
	}); err != nil {
		println(err)
		os.Exit(1)
	}
}
