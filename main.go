package main

import (
	"fmt"
	"strings"
	"unicode"
)

func translate(input string) (string, error) {

	splitline := strings.Split(input, " ")

	var piglatin []string

	for _, word := range splitline {
		if unicode.IsUpper(rune(word[0])) {
			piglatin = append(piglatin, fmt.Sprintf("%s%s%say", strings.ToUpper(string(word[1])), string(word[2:]), strings.ToLower(string(word[0]))))
		} else {
			piglatin = append(piglatin, fmt.Sprintf("%s%say", string(word[1:]), string(word[0])))
		}

	}

	return strings.Join(piglatin, " "), nil
}

func main() {

	input := "Hello World"

	final, err := translate(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(final)
}
