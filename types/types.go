package types

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Menu struct {
	options          []string
	defaultSelection *int
	currSelection    *int
}

func NewMenu(options []string, defaultSelection, currSelection *int) (*Menu, error) {
	if defaultSelection != nil {
		if *defaultSelection >= len(options) || *defaultSelection < 0 {
			return nil, fmt.Errorf("default selection should be between 0 and %d", len(options))
		}
	}
	return &Menu{
		options:          options,
		defaultSelection: defaultSelection,
	}, nil
}

func (m *Menu) GetUserInput() int {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	// 2. Ensure terminal state restores when the program exits
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	renderMenuSelection(m.options, 0, 0)
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
				fmt.Printf("\033[%dA", opLen+1)
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
				renderMenuSelection(m.options, currSlection, opLen)

			case 66:
				currSlection = (currSlection + 1) % opLen
				renderMenuSelection(m.options, currSlection, opLen)
				// case 67:
				// 	fmt.Print("\r\nRight\r\n")
				// case 68:
				// 	fmt.Print("\r\nLeft\r\n")
			}
		}
	}
	return -1
}

func (m *Menu) GetSelection(selected int) string {
	return m.options[selected]
}

func renderMenuSelection(op []string, selected, overWrite int) {
	if overWrite > 0 {
		fmt.Printf("\033[%dA", overWrite)
	}
	for i := range op {
		if i == selected {
			fmt.Printf(" \x1b[1;93m> %s\x1b[0m\r\n", op[i])
		} else {
			fmt.Printf("   %s\r\n", op[i])
		}
	}

}
