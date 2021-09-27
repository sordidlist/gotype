package target_language_keyboard_maps

import "unicode"

type RussianKeyboardMapType struct{}

func (RussianKeyboardMapType) GetLanguageName() string {
	return "Russian"
}

func (RussianKeyboardMapType) GetUppercaseKeyboardMap() map[rune]rune {
	return map[rune]rune{
		'Q':  'Й',
		'W':  'ц',
		'E':  'У',
		'R':  'К',
		'T':  'Е',
		'Y':  'Н',
		'U':  'Г',
		'I':  'Ш',
		'O':  'Щ',
		'P':  'З',
		'{':  'Х',
		'}':  'Ъ',
		'A':  'Ф',
		'S':  'Ы',
		'D':  'В',
		'F':  'А',
		'G':  'П',
		'H':  'Р',
		'J':  'О',
		'K':  'Л',
		'L':  'Д',
		';':  'Ж',
		'\'': 'Э',
		'Z':  'Я',
		'X':  'Ч',
		'C':  'С',
		'V':  'М',
		'B':  'И',
		'N':  'Т',
		'M':  'Ь',
		',':  'Б',
		'.':  'Ю',
		'/':  ',',
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
		'a': 'ф',
		's': 'ы',
		'd': 'в',
		'f': 'а',
		'g': 'п',
		'h': 'р',
		'j': 'о',
		'k': 'л',
		'l': 'д',
		';': 'ж',
		'z': 'я',
		'x': 'ч',
		'c': 'с',
		'v': 'м',
		'b': 'и',
		'n': 'т',
		'm': 'ь',
		',': 'б',
		'.': 'ю',
		'/': '.',
	}
}

func (RussianKeyboardMapType) GetShiftNumberRowMap() map[rune]rune {
	return map[rune]rune{
		'!': '!',
		'@': '"',
		'#': '№',
		'$': ';',
		'^': ':',
		'&': '?',
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
