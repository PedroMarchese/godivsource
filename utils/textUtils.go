package utils

import (
	"strings"

	"github.com/fatih/color"
)

func DivBar() {
	whiteC := color.New(color.FgWhite)
	bar := strings.Repeat("■", 78)

	whiteC.Println(bar)
}
