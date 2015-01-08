/*
Author: Harihar Subramanyam
Description:

Vigenere cipher implementation. Provide a message of lowercase letters (ex. "hello") and
a key (ex. 312) and the program will shift characters of the string to produce an encoded
string.

"h" (shift 3) = "h" -> "i" -> "j" -> "k" = "k"
"e" (shift 1) = "e" -> "f"
"l" (shift 2) = "l" -> "m" -> "n"
"l" (shift 3) = "l" -> "m" -> "n" -> "o"
"o" (shift 1) = "o" -> "p"

"kfnop"

I use this technique (with some modifications) to generate passwords that I can recompute
in my head.

Usage: go run encode.go <message> <key>
OR, (if compiled) ./encode.go <message> <key>

<message> is a string of lowercase letters.
<key> is an integer.
*/

// Package main is the default package containing the encoder.
package main

import (
	"errors"       // For making error if message isn't all lowercase letters.
	"fmt"          // For printing encoded result, usage, errors, and for converting ints to strings.
	"math"         // For computing number of digits in key's digit string.
	"os"           // For extracting command line args.
	"strconv"      // For parsing integer key.
	"unicode"      // For ensuring message is all lowercase.
	"unicode/utf8" // For convering "a" into a rune.
)

// usage prints to stdout a string explaining how to use this program.
func usage() {
	fmt.Println("Encode a message (string of lowercase letters) using a key (integer)")
	fmt.Println("Usage: go run encode.go <message> <key>")
	fmt.Println("OR (if compiled) ./encode.go <message> <key>")
}

// encode applies the Vigenere cipher to a message using the given key. It will return
// the encoded message (if there was no error) and any error that occured (or nil if no error
// ocurred).
func encode(key int, message string) (string, error) {
	result := ""

	keyDigits := toDigitSlice(key)

	// The for range loop over a string decodes runes
	for i, runeValue := range message {
		if !unicode.IsLower(runeValue) {
			return "", errors.New(fmt.Sprintf("%q is not a lowercase letter", runeValue))
		} else {
			// Encode the next letter.
			result += string(rotate(runeValue, keyDigits[i%len(keyDigits)]))
		}
	}
	return result, nil
}

// rotate shifts the given letter a specified rotation (i.e. number of letter spots). It returns
// the rotated letter. For example, rotate("a", 2) gives "c" (because rotating "a" twice gives
// "a" -> "b" -> "c", where the arrows denote rotation).
func rotate(letter rune, rotation int) rune {
	lowercaseA, _ := utf8.DecodeRuneInString("a")
	offsetFromLowercaseA := int(letter) - int(lowercaseA)
	return rune((offsetFromLowercaseA+rotation)%26 + int(lowercaseA))
}

// toDigitSlice converts an integer into a slice where each element is a digit of the integer
// (the digits appear in order, of course).
func toDigitSlice(key int) []int {
	// Make a slice with 0 length and log_10(key) + 1 capacity.
	digitSlice := make([]int, 0, int(math.Log10(float64(key)+1)))

	// Convert key to string.
	str := fmt.Sprintf("%d", key)

	for _, runeValue := range str {
		// Ignore error on parse (it's always nil because we know string(key) is parsable as an int
		// because key is an int).
		digit, _ := strconv.ParseInt(string(runeValue), 10, 0)

		digitSlice = append(digitSlice, int(digit))
	}

	return digitSlice
}

// main parses the command line arguments, does the encoding, and prints the results.
func main() {
	// Ensure correct number of arguments.
	if len(os.Args) < 3 {
		usage()
	} else {
		// Extract arguments.
		message := os.Args[1]
		key, err := strconv.ParseInt(os.Args[2], 10, 0)

		// Ensure key is integer.
		if err != nil {
			usage()
			fmt.Println("Your <key> was not an integer")
		} else {
			// Attempt encoding.
			encoded, err := encode(int(key), message)

			// Print either error or encoding.
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(encoded)
			} // else (for failed encoding check)
		} // else (for failed integer parse check)
	} // else (for # command line args check)
} // main
