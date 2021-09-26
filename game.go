package main

import (
	target_language_keyboard_maps "gotype/target-language-keyboard-maps"
)

type Game struct {
	userConfig         UserConfig
	counter            int
	arabicKeyboardMap  target_language_keyboard_maps.ArabicKeyboardMapType
	russianKeyboardMap target_language_keyboard_maps.RussianKeyboardMapType
	currentCharacters  []MultiChar
}
