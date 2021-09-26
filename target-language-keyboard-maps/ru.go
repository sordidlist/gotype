package target_language_keyboard_maps

import "unicode"

type RussianKeyboardMapType struct{}

func (RussianKeyboardMapType) GetLanguageName() string {
	return "Russian"
}

func (RussianKeyboardMapType) GetUppercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'Q': 'Й',
		'W': 'ц',
		'E': 'У',
		'R': 'К',
		'T': 'Е',
		'Y': 'Н',
		'U': 'Г',
		'I': 'Ш',
		'O': 'Щ',
		'P': 'З',
		'{': 'Х',
		'}': 'Ъ',
	}
}

func (RussianKeyboardMapType) GetLowercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'q': 'й',
		'w': 'ц',
		'e': 'у',
		'r': 'к',
		't': 'е',
		'y': 'н',
		'u': 'г',
		'i': 'ш',
		'o': 'щ',
		'p': 'з',
		'[': 'х',
		']': 'ъ',
	}
}

func (russianKeyboardMapType RussianKeyboardMapType) GetMapLength() int {
	return len(russianKeyboardMapType.GetUppercaseKeyboardMap())
}

func (russianKeyboardMapType RussianKeyboardMapType) GetUppercaseAndLowercaseKeyboardMaps() (map[rune]rune, map[rune]rune) {
	arabicMapUppercase := russianKeyboardMapType.GetUppercaseKeyboardMap()
	arabicMapLowercase := russianKeyboardMapType.GetLowercaseKeyboardMap()
	return arabicMapUppercase, arabicMapLowercase
}

func (russianKeyboardMapType RussianKeyboardMapType) GetMappedCharacter(inputChar rune) rune {
	if unicode.IsLower(inputChar) {
		kbMap := russianKeyboardMapType.GetLowercaseKeyboardMap()
		mappedChar := kbMap[inputChar]
		return mappedChar
	} else {
		kbMap := russianKeyboardMapType.GetUppercaseKeyboardMap()
		mappedChar := kbMap[inputChar]
		return mappedChar
	}
}
