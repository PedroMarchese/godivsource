package utils

import (
	"github.com/fatih/color"
)

func ErrorColor() *color.Color {
	return color.New(color.FgRed)
}

func SuccessColor() *color.Color {
	return color.New(color.FgGreen)
}

func WarningColor() *color.Color {
	return color.New(color.FgYellow)
}

func GetAllColors() (*color.Color, *color.Color, *color.Color) {
	return color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgYellow)
}
