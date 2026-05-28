package main

import (
	"strings"
)

func StringToArt(input string) string {

	if input == "" {
		return ""
	}
	//var result string
	var build strings.Builder

	var mapper = map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},

		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    /",
			"  _/ ",
			" /   ",
			"/___ ",
		},
	}

	splitInput := strings.Split(input, "\n")
	
	for _, word := range splitInput {

		for i := 0; i < 5; i++ {
			for _, ch := range word {
				if ch < '0' || ch > '9' {
					return ""
				}
				if ch >= '0' && ch <= '9' {
					build.WriteString(mapper[ch][i])
				}
			}
			build.WriteString("\n")
		}
	}
	return build.String()
}
