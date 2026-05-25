package main

import (
	"fmt"
	"strings"
)

func camelToSnake(input string) string {

	var result []string

	if input == "" {
		return ""
	}

	for i := 0; i < len(input); i++ {
		// if ch >= 'a' || ch <= 'z' {
		// 	return input
		// }
		// if ch >= 'A' || ch <= 'Z' {
		// 	return input
		// }

		if input[i] >= 'a' && input[i] <= 'z' || input[i] >= 'A' && input[i] <= 'Z' {

			if i+1 < len(input) {
				if input[i+1] >= 'A' || input[i+1] <= 'Z' {
					result = append(result, string(input[i]) + "_")
				}
			}
			result = append(result, string(input[i]))

		}
	}
	return strings.Join(result, "")
}

func main() {

	fmt.Println(camelToSnake("helloHello"))
}
