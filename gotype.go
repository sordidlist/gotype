package main

import (
	"fmt"
	"gotype/configure"
	"gotype/target-language-keyboard-maps"
)

// Look for user configs, load if they don't exist
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

type keyboardMap interface {
	GetUppercaseAndLowercaseKeyboardMaps() (map[string]string, map[string]string)
}

func main() {
	configure.GetInitialConfigs()
	ShowMainMenu()
	ar := target_language_keyboard_maps.ArabicKeyboardMapType{}

	ru := target_language_keyboard_maps.RussianKeyboardMapType{}

	printKeyboardMap(ar)
	printKeyboardMap(ru)
}

func printKeyboardMap(k keyboardMap) {
	kbmapU, kbmapL := k.GetUppercaseAndLowercaseKeyboardMaps()
	fmt.Println(kbmapU, kbmapL)
}
