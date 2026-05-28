package main

func GenerateFont() map[rune][]string {

	mapper := make(map[rune][]string)

	for r := ' '; r <= '~'; r++ {

		if r == ' ' {

			mapper[r] = []string{
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
			}
			continue
		}

		if isVowel(r) {
			mapper[r] = []string{
				"........",
				"********",
				"........",
				"********",
				"........",
				"********",
				"........",
				"********",
			}
			continue
		}

		mapper[r] = []string{
			"********",
			"........",
			"********",
			"........",
			"********",
			"........",
			"********",
			"........",
		}

	}
	return mapper
}

func isVowel(r rune) bool {

	return r == 'E' || r == 'A' || r == 'I' || r == 'O' || r == 'U'
}
