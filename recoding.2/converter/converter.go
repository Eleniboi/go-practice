package converter

import "strings"

func StringToArt(input string) string {

	if input == ""{
		return ""
	}
	var build strings.Builder
	mapper := map[rune][]string{

		'0': {

			" ___",
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
			"_____",
			"    /",
			" __/ ",
			"/    ",
			"|____",
		},
		'3': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'4': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'5': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'6': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'7': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'8': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'9': {

			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
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
