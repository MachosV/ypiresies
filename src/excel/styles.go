package excel

import "github.com/tealeg/xlsx"

var styles map[string]*xlsx.Style

func init() {
	styles = make(map[string]*xlsx.Style)
	InitStyles()
}

func InitStyles() {
	styleGrey := xlsx.NewStyle()
	styleGrey.Border.Top = "thin"
	styleGrey.Border.Bottom = "thin"
	styleGrey.Border.Right = "thin"
	styleGrey.Border.Left = "thin"
	styleGrey.Font.Name = "Verdana"
	styleGrey.Font.Size = 12
	styleGrey.Font.Bold = false
	fill := *xlsx.NewFill("solid", "FFD8D8D8", "00000000")
	styleGrey.Fill = fill
	styles["grey"] = styleGrey

}

func GetStyle(style string) *xlsx.Style {
	return styles[style]
}
