package generatefont

func GenerateFont() map[rune][]string {

	mapper := make(map[rune][]string)

	for i := ' '; i <= '~'; i++ {

		if i == ' ' {

			mapper[i] = []string{

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
		if isVowel(i) {
			mapper[i] = []string{
				"........",
				"********",
				"........",
				"********",
				"........",
				"********",
				"........",
				"********",
			}
		}
		mapper[i] = []string{
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
	return r == 'A' ||
		r == 'E' ||
		r == 'I' ||
		r == 'O' ||
		r == 'U'
}
