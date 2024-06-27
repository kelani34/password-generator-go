package main

import (
	"fmt"
	"bufio"
)


func getinput(reader *bufio.Reader, prompt string) (int, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	trimmedInput := strings.TrimSpace(input)
	number, err := strconv.Atoi(trimmedInput)
	if err != nil {
		return 0, err
	}

	return number, nil
}


func main() {
	letters := []string{	
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	numbers := []string{"0", "1", "3", "4", "5", "6", "7", "8", "9"}
	specialCharacters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+", "=", "`", "~", "{", "}", "[", "]", ";", ":", "'", ",", "<", ".", ">", "/", "?", "|", "\\"}
	
	

	fmt.Println("Welcome to the password generator.")
	numbersInput, _ := getInput(reader, "How many numbers would you like your password to have?{0-9}: ")
	lettersInput, _ := getInput(reader, "How many letters would you like your password to have?{(a-z)-(A-Z)}: ")
	specialCharactersInput, _ := getInput(reader, "How many special characters would you like your password to have?{!-|}: ")	

	fmt.Println(numbersInput, lettersInput, specialCharactersInput)

	

}
