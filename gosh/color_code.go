package gosh

import "fmt"

func setColorText(cc ColorCodes) {
	fmt.Printf("%s", cc)
}

func unsetColorText() {
	fmt.Printf("%s", ResetCC)
}

func printWithColor(text string, cc ColorCodes, sel Selector) {
	fmt.Printf(" %s%s %s%s\r\n", cc, sel, text, ResetCC)
}
