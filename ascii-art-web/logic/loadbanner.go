package logic

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
//LoadBanner read the banner file which contains the ascii art, it map and stores it value as a []string
//Each Ascii art are mapped using a rune of the valid ascii characters 
func LoadBanner(banner string) (map[rune][]string, error) {

	data, err := os.ReadFile(banner)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		return nil, errors.New("incomplete txt file")
	}
//this point is where loading takes place
	text := make(map[rune][]string)
	ascii := rune(32)
	for i := 1; i < len(lines); i += 9 {

		if ascii > 126 {
			break
		}

		text[ascii] = lines[i : i+8]
		ascii++
	}
	if len(text) != 95 {
		return nil, errors.New("incomplete characters")
	}

	return text, nil

}
