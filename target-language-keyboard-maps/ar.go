package target_language_keyboard_maps

import "unicode"

type ArabicKeyboardMapType struct{}

func (arabicKeyboardMap ArabicKeyboardMapType) GetLanguageName() string {
	return "Arabic"
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetUppercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'Q': 'ض',
		'W': 'ص',
	}
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetLowercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'q': 'ض',
		'w': 'ص',
	}
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetUppercaseAndLowercaseKeyboardMaps() (map[rune]rune, map[rune]rune) {
	arabicMapUppercase := arabicKeyboardMap.GetUppercaseKeyboardMap()
	arabicMapLowercase := arabicKeyboardMap.GetLowercaseKeyboardMap()
	return arabicMapUppercase, arabicMapLowercase
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetMapLength() int {
	return len(arabicKeyboardMap.GetUppercaseKeyboardMap())
}

func (ArabicKeyboardMapType) GetMappedCharacter(inputChar rune) rune {
	if unicode.IsLower(inputChar) {
		kbMap := ArabicKeyboardMapType{}.GetLowercaseKeyboardMap()
		mappedChar := kbMap[inputChar]
		return mappedChar
	} else {
		kbMap := ArabicKeyboardMapType{}.GetUppercaseKeyboardMap()
		mappedChar := kbMap[inputChar]
		return mappedChar
	}
}
