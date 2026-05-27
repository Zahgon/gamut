package htmlgen

import (
	"image/color"
	"io"
)

var (
	header = `
	<html>
	<body>
	`
	tablecaption = `
	<h2>%s</h2>
	`
	tableheader = `
	<table width="100%" height="64">
	<tr>
	`
	cell        = `<td bgcolor="%s" align="center"><font face="Courier" color="%s">%s</font></td>`
	tablefooter = `
	</tr>
	</table>
	`
	footer = `
	</body>
	</html>
	`
)

// Header writes the HTML header to buffer
func Header(buffer io.Writer) { _ = "STUB: not implemented"; return }

// Footer writes the HTML footer to buffer
func Footer(buffer io.Writer) { _ = "STUB: not implemented"; return }

// Cell writes a colored HTML table cell to buffer
func Cell(buffer io.Writer, c color.Color) { _ = "STUB: not implemented"; return }

// Table writes a palette of colors as an HTML table to buffer
func Table(buffer io.Writer, name string, cc []color.Color) { _ = "STUB: not implemented"; return }
