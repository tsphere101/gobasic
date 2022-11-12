package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputFloat(prompt string) (float64, error) {
	input, err := GetInput(prompt)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(input, 64)
}

func GetInputInt(prompt string) (int, error) {
	input, err := GetInput(prompt)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input)
}

func GetInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}
