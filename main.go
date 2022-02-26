package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	file = os.Stdin
	defer file.Close()

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		fmt.Println(">", inputScanner.Text())
	}
}
