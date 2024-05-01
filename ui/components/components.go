package components

import (
	"fmt"
	"himitsu/config"
	"himitsu/totp"
	"himitsu/ui/styling"
	"math"
	"strings"
	"time"

	"github.com/muesli/termenv"
)

func VersionNumber() string {
	return "   himitsu " + styling.ColorFg(config.Version, styling.HighlightedColor)
}

func Checkbox(label string, checked bool, selected bool) string {
	s := fmt.Sprintf("[ ] %s", label)
	if checked {
		s = "[x] " + label
	}
	if selected {
		return styling.ColorFg(s, styling.HighlightedColor)
	}
	return s
}

func Progressbar(width int, percent float64) string {
	w := float64(width)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += termenv.String(styling.ProgressFullChar).Foreground(styling.Term.Color(styling.Ramp[i])).String()
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(styling.ProgressEmpty, emptySize)

	return fmt.Sprintf("%s%s %02ds", fullCells, emptyCells, 30-int(math.Round(percent*30)))
}

func TOTP(currTotp totp.TOTP) string {
	code := totp.GetCode(currTotp.Secret)
	s := fmt.Sprintf("  ╔═════════════════════════════════════════════════════════════════════╗\n  ║ %s  %-25.25s  %s %s  %s ║\n  ╚═════════════════════════════════════════════════════════════════════╝\n", styling.ColorFg(fmt.Sprintf("%-15.15s", currTotp.Label), styling.HighlightedColor), currTotp.Account, code[0:3], code[3:6], Progressbar(10, float64(time.Now().UnixMilli()%30000)/30000.))
	return s
}

func KeybindsHints(keybinds []string) string {
	s := "   "
	for index, keybind := range keybinds {
		if index != 0 {
			s += styling.Dot
		}
		s += styling.Subtle(keybind)
	}
	return s
}
