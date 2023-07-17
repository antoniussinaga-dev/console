package color

import (
	"strings"

	"github.com/antoniussinaga-dev/console/types"
)

type Text string

func GetBackgroundColorCode(color string) string {
	var result string = ""
	switch color {
	case string(types.White):
		result = string(types.BackgroundWhite)
	case string(types.Red):
		result = string(types.BackgroundRed)
	case string(types.Green):
		result = string(types.BackgroundGreen)
	case string(types.Yellow):
		result = string(types.BackgroundYellow)
	case string(types.Blue):
		result = string(types.BackgroundBlue)
	case string(types.Black):
		result = string(types.BackgroundBlack)
	case string(types.Brown):
		result = string(types.BackgroundBrown)
	case string(types.Azure):
		result = string(types.BackgroundAzure)
	case string(types.Ivory):
		result = string(types.BackgroundIvory)
	case string(types.Teal):
		result = string(types.BackgroundTeal)
	case string(types.Purple):
		result = string(types.BackgroundPurple)
	case string(types.NavyBlue):
		result = string(types.BackgroundNavyBlue)
	case string(types.PeaGreen):
		result = string(types.BackgroundPeaGreen)
	case string(types.Gray):
		result = string(types.BackgroundGray)
	case string(types.Maroon):
		result = string(types.BackgroundMaroon)
	case string(types.Charcoal):
		result = string(types.BackgroundCharcoal)
	case string(types.Aquamarine):
		result = string(types.BackgroundAquamarine)
	case string(types.Coral):
		result = string(types.BackgroundCoral)
	case string(types.Fuchsia):
		result = string(types.BackgroundFuchsia)
	case string(types.Wheat):
		result = string(types.BackgroundWheat)
	case string(types.Lime):
		result = string(types.BackgroundLime)
	case string(types.Khaki):
		result = string(types.BackgroundKhaki)
	case string(types.HotPink):
		result = string(types.BackgroundHotPink)
	case string(types.Magenta):
		result = string(types.BackgroundMagenta)
	case string(types.Olden):
		result = string(types.BackgroundWhite)
	case string(types.Plum):
		result = string(types.BackgroundPlum)
	case string(types.Olive):
		result = string(types.BackgroundOlive)
	case string(types.Cyan):
		result = string(types.BackgroundCyan)
	default:
		result = ""
	}
	return result
}

func GetFontStyle(style string) string {
	var result string = ""
	switch strings.ToLower(style) {
	case types.Bold:
		result = types.StyleBold
	default:
		result = ""
	}
	return result
}
