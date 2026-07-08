package main

import (
	"log"

	"github.com/naghulvarshan/gosh"
)

func main() {
	colorMenu, _ := gosh.NewMenu(
		[]gosh.Options{
			{DisplayName: "Yellow", Linker: gosh.Yellow},
			{DisplayName: "Black", Linker: gosh.Black},
			{DisplayName: "Blue", Linker: gosh.Blue},
			{DisplayName: "Green", Linker: gosh.Green},
			{DisplayName: "Red", Linker: gosh.Red}},
		nil)
	col := colorMenu.GetSelection(colorMenu.GetUserInput())
	op := col.Linker.(gosh.ColorCodes)
	menu, _ := gosh.NewMenu(
		[]gosh.Options{
			{DisplayName: "1. Apple"},
			{DisplayName: "2. Oranges"},
			{DisplayName: "3. Bananas"},
		}, &op)
	inp := menu.GetUserInput()
	if inp == -1 {
		log.Println("No input selected")
	} else {
		log.Println("user selection is: ", menu.GetSelection(inp).DisplayName)
	}
}
