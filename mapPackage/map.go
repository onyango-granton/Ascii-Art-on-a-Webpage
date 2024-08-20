package mapPackage

import (
	"ascii-web-art/utils"
	"fmt"
	"os"
	"strings"
)

// AsciiMapping given a banner file, reads all graphics representations of the ASCII characters and
// returns a map of the ASCII character to the graphics representations of the ASCII character
func AsciiMapping(patternFile string) map[rune][]string {
	var splitted []string

	if patternFile == "thinkertoy.txt" {
		testfile, err1 := os.ReadFile(patternFile)
		if !(utils.CheckFileIntegrity(testfile)){
			fmt.Fprintln(os.Stderr, "error:", patternFile, "has been modified")
			return map[rune][]string{}
		}
		if len(testfile) == 0 {
			fmt.Fprintln(os.Stderr, "error:", patternFile, "is empty or doesn't exist ")
			return map[rune][]string{}
		} else if err1 != nil {
			fmt.Println(err1.Error())
			return map[rune][]string{}
		}

		splitted = strings.Split(string(testfile), "\r\n") // strings of thinkeratoi are seperated by \r\n [13,10]
	} else {
		testfile, err := os.ReadFile(patternFile)
		if !(utils.CheckFileIntegrity(testfile)){
			fmt.Fprintln(os.Stderr, "error:", patternFile, "has been modified")
			return map[rune][]string{}
		}
		if len(testfile) == 0 {
			fmt.Println("error:", patternFile, "is empty")
			return map[rune][]string{}
		}
		if err != nil {
			fmt.Println(err.Error())
			return map[rune][]string{}
		}

		splitted = strings.Split(string(testfile), "\n")
	}

	asciiMapping := make(map[rune][]string)
	startAscii := ' '
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciiMapping[startAscii] = arrayString
		startAscii++
	}

	return asciiMapping
}
