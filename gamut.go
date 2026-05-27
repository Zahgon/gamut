package gamut

import (
	"image/color"
)

// A Palette is a collection of colors
type Palette struct {
	colors map[color.Color]Colors
	names  map[string]color.Color
}

// MixedWith mixes two palettes
func (g Palette) MixedWith(p Palette) Palette { _ = "STUB: not implemented"; return *new(Palette) }

// AddColors adds colors to the palette
func (g *Palette) AddColors(cc Colors) { _ = "STUB: not implemented"; return }

// Colors returns the Palette's colors
func (g Palette) Colors() Colors { _ = "STUB: not implemented"; return *new(Colors) }

// Clamped expects a slice of colors and returns a slice of the nearest matching
// colors from the palette
func (g Palette) Clamped(cc []color.Color) Colors { _ = "STUB: not implemented"; return *new(Colors) }

// Color returns the color with a specific name
func (g Palette) Color(name string) (color.Color, bool) {
	_ = "STUB: not implemented"
	return *new(color.Color), false
}

// Name returns the name of the closest matching color
func (g Palette) Name(color color.Color) (Colors, float64) {
	_ = "STUB: not implemented"
	return *new(Colors), 0
}

// Filter returns colors matching name
func (g Palette) Filter(name string) Colors { _ = "STUB: not implemented"; return *new(Colors) }
