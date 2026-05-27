package palette

import (
	"github.com/muesli/gamut"
)

var (
	// Wikipedia Palette
	Wikipedia gamut.Palette
	// Monokai Palette
	Monokai gamut.Palette
	// Crayola Palette
	Crayola gamut.Palette
	// Resene Palette
	Resene gamut.Palette
	// RAL Palette
	RAL gamut.Palette
	// CSS Named Color Palette
	CSS gamut.Palette
)

// AllPalettes returns a mix of all palettes defined in gamut
func AllPalettes() gamut.Palette { _ = "STUB: not implemented"; return *new(gamut.Palette) }
