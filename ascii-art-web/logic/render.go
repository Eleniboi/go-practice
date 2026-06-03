package logic

import "strings"
//RenderLine gives the input to the already mapped ascii art
//Each ascii art contain 8 rows, RenderLine return 8 rows of Ascii art for each character in the input string
func RenderLine(input string, banner map[rune][]string) []string {

	var result []string

	var build strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range input {
			build.WriteString(banner[char][i])
		}
		result = append(result, build.String())
		build.Reset()
	}
	return result
}
