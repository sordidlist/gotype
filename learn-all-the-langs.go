package main

import (
	"fmt"
	"learn-all-the-langs/configure"
	"learn-all-the-langs/target-language-keyboard-maps"
	"learn-all-the-langs/views"
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

func main() {
	configure.GetInitialConfigs()
	views.ShowMainMenu()
	ar := target_language_keyboard_maps.ArabicKeyboardMapType{}
	arMapUpper, arMapLower := ar.GetUppercaseAndLowercaseKeyboardMaps()
	fmt.Printf(arMapUpper["E"])
	fmt.Printf(arMapLower["u"])
}
