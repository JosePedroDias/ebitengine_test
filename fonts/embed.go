package fonts

import (
	_ "embed"
)

var (
	// from here: https://www.fontsquirrel.com/fonts/Silkscreen

	//go:embed slkscr.ttf
	Silkscreen_regular []byte

	//go:embed slkscrb.ttf
	Silkscreen_bold []byte
)
