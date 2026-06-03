package logic

import "fmt"

//ValidateInput checks for supported ascii character, it returns a string and an error
func ValidateInput(input string) (string, error) {

	for _, ch := range input {

		if ch < 32 || ch > 126 {

			return "", fmt.Errorf("%q is not a supported ascii character", ch)
		}
	}
	return input, nil
}
