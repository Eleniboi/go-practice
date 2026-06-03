package logic


import (

	"strings"
)

//GenerateArt returns the final ascii art, it is just like the brain of the project it connect other function and form the ascii art base on the input.
//it make the work less for main, the main function just have to execute it.
func GenerateArt(input string, banner map[rune][]string) string {

	var build strings.Builder

	if input == ""{
		return ""
	}

	slitinput := SplitInput(input)
	allempty := true

	for _, word := range slitinput{
		if word != "" {
			allempty = false
		}
	}

	if allempty{

		for i := 0; i < len(slitinput)-1;  i++{
			build.WriteString("\n")
		}
		return build.String()
	}

	for _, word := range slitinput{

		if word == ""{
			build.WriteString("\n")
			continue
		}
//this point render function is called to perform it job, which is rendering input to the banner.
		result := RenderLine(word, banner)

		for i := 0; i < len(result); i++{

			build.WriteString(result[i])
			build.WriteString("\n")
		}
	}
	return build.String()

}