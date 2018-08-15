package excel

import (
	"github.com/tealeg/xlsx"
)

func GetStyle(color string) *xlsx.Style {
	var style *xlsx.Style
	style = xlsx.NewStyle()
	style.Font.Name = "Verdana"
	style.Font.Size = 12
	style.Font.Bold = false
	fill := xlsx.NewFill("solid", color, "00000000")
	style.Fill = *fill
	return style
}
