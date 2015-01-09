/*
Author: Harihar Subramanyam
Description:

Command line program for applying the vigenere cipher
*/

// Package main is the default package containing the encoder.
package main

import (
	"fmt"                                                     // For printing encoded result, usage, errors, and for converting ints to strings.
	"github.com/hariharsubramanyam/password_encoder/vigenere" // For encoding.
	"os"                                                      // For extracting command line args.
	"strconv"                                                 // For parsing integer key.
)

// usage prints to stdout a string explaining how to use this program.
func usage() {
	fmt.Println("Encode a message (string of lowercase letters) using a key (integer)")
	fmt.Println("Usage (if installed): password_encoder <message> <key>")
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
			encoded, err := vigenere.Encode(int(key), message)

			// Print either error or encoding.
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(encoded)
			} // else (for failed encoding check)
		} // else (for failed integer parse check)
	} // else (for # command line args check)
} // main
