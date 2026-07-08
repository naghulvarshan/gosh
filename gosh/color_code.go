package gosh

import "fmt"

func setColorText(cc ColorCodes) {
	fmt.Printf("%s", cc)
}

func unsetColorText() {
	fmt.Printf("%s", ResetCC)
}

func printWithColor(text string, cc ColorCodes) {
	fmt.Printf(" %s> %s%s\r\n", cc, text, ResetCC)
}
