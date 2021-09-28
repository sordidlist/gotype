package target_language_keyboard_maps

import "unicode"

type ArabicKeyboardMapType struct{}

func (arabicKeyboardMap ArabicKeyboardMapType) GetLanguageName() string {
	return "Arabic"
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetUppercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'Q': 'َ',
		'W': 'ً',
		'E': 'ُ',
		'R': 'ٌ',
		'T': 'ﻹ',
		'Y': 'إ',
		'U': '`',
		'I': '÷',
		'O': '×',
		'P': '؛',
		'{': '<',
		'}': '>',
		'|': '…',
		'A': 'ِ',
		'S': 'ٍ',
		'D': ']',
		'F': '[',
		'G': 'ﻷ',
		'H': 'أ',
		'J': 'ـ',
		'K': '،',
		'L': '/',
		'Z': '~',
		'X': 'ْ',
		'C': '}',
		'V': '{',
		'B': 'ﻵ',
		'N': 'آ',
		'M': '\'',
		'<': ',',
		'>': '.',
		'?': '؟',
	}
}

func (arabicKeyboardMap ArabicKeyboardMapType) GetLowercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'q': 'ض',
		'w': 'ص',
		'e': 'ث',
		'r': 'ق',
		't': 'ف',
		'y': 'غ',
		'u': 'ع',
		'i': 'ه',
		'o': 'خ',
		'p': 'ح',
		'[': 'ج',
		']': 'د',
		'a': 'ش',
		's': 'س',
		'd': 'ي',
		'f': 'ب',
		'g': 'ل',
		'h': 'ا',
		'j': 'ت',
		'k': 'ن',
		'l': 'م',
		'z': 'ئ',
		'x': 'ء',
		'c': 'ؤ',
		'v': 'ر',
		'b': 'ﻻ',
		'n': 'ى',
		'm': 'ة',
		',': 'و',
		'.': 'ز',
		'/': 'ظ',
	}
}

func (ArabicKeyboardMapType) GetShiftNumberRowMap() map[rune]rune {
	return map[rune]rune{}
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
