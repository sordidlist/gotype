package target_language_keyboard_maps

type keyboardMap interface {
	getUppercaseAndLowercaseKeyboardMaps() (map[string]string, map[string]string)
}
