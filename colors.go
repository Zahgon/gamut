package gamut

import (
	"image/color"
)

// A Color is a color including its name and reference URL
type Color struct {
	Name      string
	Color     color.Color
	Reference string
}

// Colors is a slice of colors
type Colors []Color

// Hex returns the color encoded by a hex-string, e.g. "#ABCDEF".
func Hex(s string) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

// ToHex returns the hex encoding of a color, e.g. "#ABCDEF".
func ToHex(c color.Color) string { _ = "STUB: not implemented"; return "" }

// HueOffset returns color with a different hue angle
func HueOffset(c color.Color, degrees int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

// Tetradic returns the tetradic values for any given color
func Tetradic(c1 color.Color, c2 color.Color) []color.Color { _ = "STUB: not implemented"; return nil }

// Triadic returns the triadic values for any given color
func Triadic(c color.Color) []color.Color { _ = "STUB: not implemented"; return nil }

// Quadratic returns the quadratic values for any given color
func Quadratic(c color.Color) []color.Color { _ = "STUB: not implemented"; return nil }

// Analogous returns the analogous values for any given color
func Analogous(c color.Color) []color.Color { _ = "STUB: not implemented"; return nil }

// SplitComplementary returns the split complementary values for any given color
func SplitComplementary(c color.Color) []color.Color { _ = "STUB: not implemented"; return nil }

// Complementary returns the complementary value for any given color
func Complementary(c color.Color) color.Color {
	_ = "STUB: not implemented"
	return *

	// Contrast returns the color with the most contrast (hence either black or white)
	new(color.Color)
}

func Contrast(c color.Color) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

// Monochromatic returns the specified amount of monochromatic colors based on
// a given color's hues
func Monochromatic(c color.Color, count int) []color.Color { _ = "STUB: not implemented"; return nil }

// Blends returns a slice of interpolated colors, blended between two colors
func Blends(c1, c2 color.Color, count int) []color.Color { _ = "STUB: not implemented"; return nil }

// Shades returns the specified amount of a color's shades
func Shades(c color.Color, count int) []color.Color { _ = "STUB: not implemented"; return nil }

// Tints returns the specified amount of a color's tints
func Tints(c color.Color, count int) []color.Color { _ = "STUB: not implemented"; return nil }

// Tones returns the specified amount of a color's tone
func Tones(c color.Color, count int) []color.Color { _ = "STUB: not implemented"; return nil }

// Cool returns whether a color is considered to have a cool temperature
func Cool(c color.Color) bool { _ = "STUB: not implemented"; return false }

// Warm returns whether a color is considered to have a warm temperature
func Warm(c color.Color) bool {
	_ = "STUB: not implemented"

	// Lighter returns a lighter version of the specified color
	return false
}

func Lighter(c color.Color, percent float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

// Darker returns a darker version of the specified color
func Darker(c color.Color, percent float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}
