package main

import "fmt"

func main() {
	letters := []string{	
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	numbers := []string{"0", "1", "3", "4", "5", "6", "7", "8", "9"}
	specialCharacters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+", "=", "`", "~", "{", "}", "[", "]", ";", ":", "'", ",", "<", ".", ">", "/", "?", "|", "\\"}
	
	fmt.Println(letters, numbers, specialCharacters)
}
