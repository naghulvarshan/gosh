package main

import (
	"log"

	"github.com/naghulvarshan/gosh"
)

func main() {
	colorMenu := gosh.NewMenu(
		[]gosh.Options{
			{DisplayName: "Yellow", Linker: gosh.Yellow, CustomSelectionColor: &gosh.Yellow},
			{DisplayName: "Black", Linker: gosh.Black, CustomSelectionColor: &gosh.Black},
			{DisplayName: "Blue", Linker: gosh.Blue, CustomSelectionColor: &gosh.Blue},
			{DisplayName: "Green", Linker: gosh.Green, CustomSelectionColor: &gosh.Green},
			{DisplayName: "Red", Linker: gosh.Red, CustomSelectionColor: &gosh.Red}},
	)
	col := colorMenu.GetSelection(colorMenu.GetUserInput())
	op := col.Linker.(gosh.ColorCodes)
	selectorMenu := gosh.NewMenu(
		[]gosh.Options{
			{DisplayName: "Star", Linker: gosh.StarSelector},
			{DisplayName: "GT", Linker: gosh.GTSelector},
			{DisplayName: "Bullet", Linker: gosh.BulletSelector}},
	)
	selector := selectorMenu.GetSelection(selectorMenu.GetUserInput()).Linker.(gosh.Selector)
	menu := gosh.NewMenu(
		[]gosh.Options{
			{DisplayName: "1. Apple"},
			{DisplayName: "2. Oranges"},
			{DisplayName: "3. Bananas"},
		}, gosh.WithSelectColor(op), gosh.WithSelector(selector))
	inp := menu.GetUserInput()
	if inp == -1 {
		log.Println("No input selected")
	} else {
		log.Println("user selection is: ", menu.GetSelection(inp).DisplayName)
	}
}
