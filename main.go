package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

	password := []string{}	

	// randomize passowrd
	for i := 0; i <= numbersInput; i++ {
		randomIndex := rand.Intn(len(numbers))
		password = append(password, numbers[randomIndex])
	}
	for i := 0; i <= lettersInput; i++ {
		randomIndex := rand.Intn(len(letters))
		password = append(password, letters[randomIndex])
	}
	for i := 0; i <= specialCharactersInput; i++ {
		randomIndex := rand.Intn(len(specialCharacters))
		password = append(password, specialCharacters[randomIndex])
	}	

	// shuffle password

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	passwordResult := strings.Join(password, "")
	fmt.Println(passwordResult)


}
