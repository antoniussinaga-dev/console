package console

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/antoniussinaga-dev/console/color"
	"github.com/antoniussinaga-dev/console/types"
	"github.com/antoniussinaga-dev/console/utils"
)

type ColorOptions types.ColorOptions
type FontStyle string

func Log(n ...interface{}) {
	if len(n) == 0 {
		fmt.Println("\033[90mNil\033[0m")
	} else {
		for _, t := range n {
			structType := reflect.TypeOf(t)
			structValue := reflect.ValueOf(t)
			if reflect.Type.String(structType) == "string" {
				fmt.Println(utils.GetDataWithType("string", structValue.Interface()))
			} else if reflect.Type.String(structType) != "string" {
				fmt.Print(FontStyle(" \033[1m\033[94m{\n\033[0m").Blue())
				for i := 0; i < structType.NumField(); i++ {
					field := structType.Field(i)
					fieldName := field.Name
					fieldType := field.Type.String()
					// fmt.Println(fieldType)
					fieldValue := structValue.Field(i).Interface()
					value := utils.GetDataWithType(fieldType, fieldValue)
					result := fmt.Sprintf("     \033[1m\033[90m%v : \033[0m%v\n", fieldName, value)
					fmt.Print(result)
				}
				fmt.Print(FontStyle(" \033[1m\033[94m}\n\033[0m").Blue())
			} else {
				fmt.Println(t)
			}
		}
	}
}

func Exit() {
	os.Exit(0)
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Version() string {
	return "1.0.0"
}
func (s FontStyle) White() FontStyle {
	return FontStyle(types.ColorWhite) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Red() FontStyle {
	return FontStyle(types.ColorRed) + s + FontStyle(types.ColorReset)
}
func (s FontStyle) Green() FontStyle {
	return FontStyle(types.ColorGreen) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Yellow() FontStyle {
	return FontStyle(types.ColorYellow) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Black() FontStyle {
	return FontStyle(types.ColorBlack) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Blue() FontStyle {
	return FontStyle(types.ColorBlue) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Brown() FontStyle {
	return FontStyle(types.ColorBrown) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Azure() FontStyle {
	return FontStyle(types.ColorAzure) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Ivory() FontStyle {
	return FontStyle(types.ColorIvory) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Teal() FontStyle {
	return FontStyle(types.ColorTeal) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Purple() FontStyle {
	return FontStyle(types.ColorPurple) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) NavyBlue() FontStyle {
	return FontStyle(types.ColorNavyBlue) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) PeaGreen() FontStyle {
	return FontStyle(types.ColorPeaGreen) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Gray() FontStyle {
	return FontStyle(types.ColorGray) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Maroon() FontStyle {
	return FontStyle(types.ColorMaroon) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Charcoal() FontStyle {
	return FontStyle(types.ColorCharcoal) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Aquamarine() FontStyle {
	return FontStyle(types.ColorAquamarine) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Coral() FontStyle {
	return FontStyle(types.ColorCoral) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Fuchsia() FontStyle {
	return FontStyle(types.ColorFuchsia) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Wheat() FontStyle {
	return FontStyle(types.ColorWheat) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Lime() FontStyle {
	return FontStyle(types.ColorLime) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Khaki() FontStyle {
	return FontStyle(types.ColorKhaki) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) HotPink() FontStyle {
	return FontStyle(types.ColorHotPink) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Olden() FontStyle {
	return FontStyle(types.ColorOlden) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Plum() FontStyle {
	return FontStyle(types.ColorPlum) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Olive() FontStyle {
	return FontStyle(types.ColorOlive) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Cyan() FontStyle {
	return FontStyle(types.ColorCyan) + s + FontStyle(types.ColorReset)
}

func (s FontStyle) Reset() FontStyle {
	return s + FontStyle(types.ColorReset)
}

// ADD PADING FUNC
func (s FontStyle) Padding(n int) FontStyle {
	padding := strings.Repeat(" ", n)
	return FontStyle(padding + string(s) + padding)
}
func (s FontStyle) PaddingLeft(n int) FontStyle {
	padding := strings.Repeat(" ", n)
	return FontStyle(padding + string(s))
}
func (s FontStyle) PaddingRight(n int) FontStyle {
	padding := strings.Repeat(" ", n)
	return FontStyle(string(s) + padding)
}

// FONT STYLE
func (s FontStyle) Bold() FontStyle {
	fst := color.GetFontStyle("bold")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Faint() FontStyle {
	fst := color.GetFontStyle("faint")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Italic() FontStyle {
	fst := color.GetFontStyle("italic")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Underline() FontStyle {
	fst := color.GetFontStyle("underline")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Blink() FontStyle {
	fst := color.GetFontStyle("blink")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Inverse() FontStyle {
	fst := color.GetFontStyle("inverse")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Conceal() FontStyle {
	fst := color.GetFontStyle("conceal")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}
func (s FontStyle) Strikethrought() FontStyle {
	fst := color.GetFontStyle("strikethrought")
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(fst + text + "\033[0m")
}

// END FONT STYLE

func (s FontStyle) BackgroundColor(n string) FontStyle {
	bg := color.GetBackgroundColorCode(strings.ToLower(n))
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(bg + text + "\033[0m")
}

// AND FUNC TEXT DECORATION
func (s FontStyle) FontStyle(td string) FontStyle {
	decoration := color.GetFontStyle(td)
	text := strings.ReplaceAll(string(s), string(types.ColorReset), "")
	return FontStyle(decoration + text + "\033[0m")
}

// ADITIONAL
func (s FontStyle) PrintLog() {
	fmt.Println(s)
}
