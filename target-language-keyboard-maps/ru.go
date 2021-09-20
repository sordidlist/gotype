package target_language_keyboard_maps

type russianKeyboardMapType struct{}

func (russianKeyboardMap russianKeyboardMapType) getUppercaseAndLowercaseKeyboardMaps() (map[string]string, map[string]string) {
	russianMapLowercase := map[string]string{
		"q": "й",
		"w": "ц",
		"e": "у",
		"r": "к",
		"t": "е",
		"y": "н",
		"u": "г",
		"i": "ш",
		"o": "щ",
		"p": "з",
	}

	russianMapUppercase := map[string]string{
		"Q": "Й",
		"W": "Ц",
		"E": "У",
		"R": "К",
		"T": "Е",
		"Y": "Н",
		"U": "Г",
		"I": "Ш",
		"O": "Щ",
		"P": "З",
	}

	return russianMapUppercase, russianMapLowercase
}
