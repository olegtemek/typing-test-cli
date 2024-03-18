package utils

import "fmt"

func ColorizePrint(format string, color string) {
	switch color {
	case "red":
		color = "\u001b[31m"
	case "bgRed":
		color = "\u001b[41m"
	case "green":
		color = "\u001b[32m"
	case "bgGreen":
		color = "\u001b[42m"
	case "yellow":
		color = "\u001b[33m"
	case "gray":
		color = "\u001b[37m"
	default:
		color = "\u001b[0m"
	}

	fmt.Printf("%s%s\033[0m", color, format)
}
