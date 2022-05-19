package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("1.Txt  2.HTML: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	input = strings.Replace(input, "\n", "", -1)

	for !validateInput(input) {
		fmt.Print("Invalid input, please try again \n")
		return GetInput()
	}
	return input, nil
}

func validateInput(input string) bool {
	if input == "1" || input == "2" {
		return true
	}
	return false
}
