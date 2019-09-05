package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	salutation := getSalutation()
	greeting := fmt.Sprintf("%s from go", salutation)
	color.Green(greeting)
}

func getSalutation() string {
	return "Hello"
}
