package main

import "fmt"

func printTextColor256(s string, color int) {
	printTextColoredBackground256(s, color)
	printTextColoredText256(s, color)
}

func printTextColoredBackground256(s string, color int) {
	str := fmt.Sprintf("\x1b[48;5;%dm\x1b[30m%3s\x1b[0m ", color, s)
	fmt.Println(str)
}

func printTextColoredText256(s string, color int) {
	str := fmt.Sprintf("\x1b[38;5;%dm%3s\x1b[0m ", color, s)
	fmt.Println(str)
}
