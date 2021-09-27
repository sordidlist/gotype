package main

type keyboardMap interface {
	GetUppercaseKeyboardMap() map[rune]rune
	GetLowercaseKeyboardMap() map[rune]rune
	GetUppercaseAndLowercaseKeyboardMaps() (map[rune]rune, map[rune]rune)
	GetShiftNumberRowMap() map[rune]rune
	GetLanguageName() string
	GetMapLength() int
}
