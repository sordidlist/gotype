package target_language_keyboard_maps

type ArabicKeyboardMapType struct{}

func (arabicKeyboardMap ArabicKeyboardMapType) GetUppercaseAndLowercaseKeyboardMaps() (map[string]string, map[string]string) {
	arabicMapUppercase := map[string]string{
		"Q": "ض",
		"W": "ص",
		"E": "ث",
		"R": "ق",
		"T": "ف",
		"Y": "غ",
		"U": "ع",
		"I": "ه",
		"O": "خ",
		"P": "ح",
	}

	arabicMapLowercase := map[string]string{
		"q": "ض",
		"w": "ص",
		"e": "ث",
		"r": "ق",
		"t": "ف",
		"y": "غ",
		"u": "ع",
		"i": "ه",
		"o": "خ",
		"p": "ح",
	}

	return arabicMapUppercase, arabicMapLowercase
}
