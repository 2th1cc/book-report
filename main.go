package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data/frankenstein.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	counts := make(map[rune]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		counts[rune(scanner.Text()[0])]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	for char, count := range counts {
		fmt.Printf("%q: %d\n", char, count)
	}
}
