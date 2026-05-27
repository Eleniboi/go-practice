package main

func GeneratePattern(c rune) []string {

	var result []string

	if c < 'A' || c > 'Z' {
		return result
	}

	if c == 'A' {

		result = []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
		return result
	}
	if c == 'Z' {
		result = []string{
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		}
		return result

	}

	if c >= 'B' || c <= 'Y' {

		result = []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
		return result
	}

	return result
}
