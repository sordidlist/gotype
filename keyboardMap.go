package main

import target_language_keyboard_maps "gotype/target-language-keyboard-maps"

type keyboardMap interface {
	GetUppercaseKeyboardMap() map[rune]rune
	GetLowercaseKeyboardMap() map[rune]rune
	GetUppercaseAndLowercaseKeyboardMaps() (map[rune]rune, map[rune]rune)
	GetLanguageName() string
	GetMapLength() int
}

func getLanguageKeyboardMap(selectedLanguage string) keyboardMap {
	if selectedLanguage == "Arabic" {
		return target_language_keyboard_maps.ArabicKeyboardMapType{}
	}
	if selectedLanguage == "Russian" {
		return target_language_keyboard_maps.RussianKeyboardMapType{}
	}
	return nil
}
