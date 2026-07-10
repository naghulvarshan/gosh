package gosh

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

// NewMenu returns a Menu with given options as the choices displayed
// Custom behavior like changing selection color, changing select icon can be set using
// MenuOpts
func NewMenu(options []Options, opts ...MenuOpts) *Menu {
	menu := &Menu{
		options:        options,
		selectionColor: Yellow,
		selector:       GTSelector,
	}
	for _, opt := range opts {
		opt(menu)
	}
	return menu
}

// Overrides the default select color.
// Select color is the color the text is displayed when that item is the current choice
func WithSelectColor(cc ColorCodes) MenuOpts {
	return func(m *Menu) {
		m.selectionColor = cc
	}
}

// Overrides the default selector
// Selector is the symbol before the item name currently selected
func WithSelector(sel Selector) MenuOpts {
	return func(m *Menu) {
		m.selector = sel
	}
}

// GetUserInput displays the menu and waits for the user to choose an option.
// The user can move between options with up or down arrows, and press Enter to choose.
// To abort the customer does Ctrl + c.
// The option index is returned. For aborted action, -1 is returned.
func (m *Menu) GetUserInput() int {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	// Ensure terminal state restores when the program exits
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	m.renderMenuSelection(0, 0)
	currSlection := 0
	opLen := len(m.options)
	buf := make([]byte, 3)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		if n == 1 {
			switch buf[0] {
			case 13, 10: // Enter
				fmt.Printf("\033[%dA", opLen)
				fmt.Print("\033[J")
				fmt.Print("\r\nEnter pressed\r\n")
				return currSlection
			case 3: // Ctrl+C
				return -1
			}
		} else if n == 3 && buf[0] == 27 && buf[1] == 91 { // ESC [
			switch buf[2] {
			case 65:
				currSlection = currSlection - 1
				if currSlection < 0 {
					currSlection = currSlection + opLen
				}
				m.renderMenuSelection(currSlection, opLen)

			case 66:
				currSlection = (currSlection + 1) % opLen
				m.renderMenuSelection(currSlection, opLen)
				// case 67:
				// 	fmt.Print("\r\nRight\r\n")
				// case 68:
				// 	fmt.Print("\r\nLeft\r\n")
			}
		}
	}
	return -1
}

func (m *Menu) GetSelection(selected int) Options {
	if selected < 0 || selected > len(m.options) {
		log.Fatalf("the selected index %d is out of bounds.", selected)
	}
	return m.options[selected]
}

func (m *Menu) renderMenuSelection(selected, overWrite int) {
	if overWrite > 0 {
		fmt.Printf("\033[%dA", overWrite)
	}
	for i := range m.options {
		if i == selected {
			printColor := m.selectionColor
			if m.options[i].CustomSelectionColor != nil {
				printColor = *m.options[i].CustomSelectionColor
			}
			printWithColor(m.options[i].DisplayName, printColor, m.selector)
		} else {
			fmt.Printf("   %s\r\n", m.options[i].DisplayName)
		}
	}
}
