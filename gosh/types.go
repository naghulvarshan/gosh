package gosh

type Menu struct {
	options        []Options
	selectionColor ColorCodes
	selector       Selector
}

type Options struct {
	DisplayName          string
	Description          *string //TODO display description when selected
	Linker               interface{}
	CustomSelectionColor *ColorCodes
}

//Color codes

type ColorCodes string

var (
	Yellow  ColorCodes = "\x1b[1;93m"
	Black   ColorCodes = "\x1b[1;90m"
	Red     ColorCodes = "\x1b[1;91m"
	Green   ColorCodes = "\x1b[1;92m"
	Blue    ColorCodes = "\x1b[1;94m"
	ResetCC ColorCodes = "\x1b[0m"
)

type Selector string

var (
	GTSelector     Selector = ">"
	StarSelector   Selector = "*"
	BulletSelector Selector = "."
)

type MenuOpts func(*Menu)
