package ascii

import (
	"strings"
)

// AsciiArt creates ASCII art for the given words based on the contents provided.
func AsciiArt(words, contents []string) string {
	var countSpace int
	var artBuild strings.Builder // Using strings.Builder for efficient string concatenation.

	sub := false // checker
	position := 0  
	subStr := "he"
	for _, word := range words {
		indexs := subIndexs(word, subStr)

		if word != "" {
			for i := 0; i < 8; i++ { // Each character is represented by 8 lines of ASCII art.
				for j, char := range word {
					if validIndex(j, indexs){
						artBuild.WriteString("\033[32m")
						sub = true
						position = j
					}
					if char == '\n' {
						continue
					}
					// Ensure the character is a printable ASCII character.
					if !(char >= 32 && char <= 126) {
						return "Error: Non printable ASCII character\n"
					}
					// Append the corresponding ASCII art line for the character.
					artBuild.WriteString(contents[int(char-' ')*9+1+i])

					if sub && j == position+len(subStr)-1{
						artBuild.WriteString("\033[37m")
					}
				}
				artBuild.WriteRune('\n')
			}
		} else { // Handle spaces between words.
			countSpace++
			if countSpace < len(words) {
				artBuild.WriteRune('\n')
			}
		}
	}
	return artBuild.String() // Return the constructed ASCII art as a string.
}


func subIndexs(s, subStr string)[]int{
    index := []int{}
    leftCharacters := 0
    
    for{
        idx:= strings.Index(s, subStr)
        if idx == -1{
            break
        }
        index = append(index, idx+leftCharacters)
        s = s[idx+len(subStr):]
        leftCharacters += idx+len(subStr)
    }
    return index
}

func validIndex(index int, indexs []int)bool{
	for _, idx := range indexs{
		if index == idx{
			return true
		}
	}
	return false
}