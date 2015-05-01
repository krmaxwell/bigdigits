package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0  ",
		"  000   "},
	{" 1 ",
		"11 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		"111"},
	{" 222  ",
		"2    2",
		"    2 ",
		"   2  ",
		"  2   ",
		" 2    ",
		"22222 "},
	{" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 "},
	{"4   4",
		"4   4",
		"4   4",
		"44444",
		"    4",
		"    4",
		"    4"},
	{"55555",
		"5    ",
		"5    ",
		" 555 ",
		"    5",
		"5   5",
		" 555 "},
	{" 666 ",
		"6    ",
		"6    ",
		"6666 ",
		"6   6",
		"6   6",
		" 666 "},
	{"7777",
		"   7",
		"  7 ",
		"  7 ",
		" 7  ",
		"7   ",
		"7   "},
	{" 888 ",
		"8   8",
		"8   8",
		" 888 ",
		"8   8",
		"8   8",
		" 888 "},
	{" 999 ",
		"9   9",
		"9   9",
		" 9999",
		"   9 ",
		"   9 ",
		"  9  "}}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] { // each number should be seven lines long
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid number, is it a whole number?")
			}
		}
		fmt.Println(line)
	}
}
