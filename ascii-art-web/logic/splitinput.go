package logic

import "strings"
//SplitInput split input by a literal "\n", and return the slice of the input
func SplitInput(input string) []string {
	return strings.Split(input, "\\n")
}
