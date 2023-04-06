package passwordgenerator

import (
	"flag"
	"log"
	"math/rand"
)

const (
	numbers       = "1234567890"
	logograms     = "#$%&@^~`"
	punctuations  = ".,:;"
	quotations    = `"'`
	dashesSlashes = `\/|_-`
	mathSymbols   = "<>*+!?="
	brackets      = "{[()]}"
)

// Function to randomize the password string
func Randomize(length int, baseString string, useNumbers, useLogograms, usePunctuations, useQuotations, useDashesSlashes, useMathSymbols, useBrackets bool) string {

	// Define variables for the new basestring and the temporary password
	newBaseString := baseString
	tempPWD := ""

	// Add the charachters from "Numbers" to the "newBaseString" variable and Add 1 random char from "numbers" to the variable "tempPWD"
	if useNumbers {
		newBaseString += numbers
		tempPWD += string(numbers[rand.Intn(len(numbers))])
	}

	// Add the charachters from "logograms" to the "newBaseString" variable and Add 1 random char from "logograms" to the variable "tempPWD"
	if useLogograms {
		newBaseString += logograms
		tempPWD += string(logograms[rand.Intn(len(logograms))])
	}

	// Add the charachters from "punctuations" to the "newBaseString" variable and Add 1 random char from "punctuations" to the variable "tempPWD"
	if usePunctuations {
		newBaseString += punctuations
		tempPWD += string(punctuations[rand.Intn(len(punctuations))])
	}

	// Add the charachters from "quotations" to the "newBaseString" variable and Add 1 random char from "quotations" to the variable "tempPWD"
	if useQuotations {
		newBaseString += quotations
		tempPWD += string(quotations[rand.Intn(len(quotations))])
	}

	// Add the charachters from "dashesSlashes" to the "newBaseString" variable and Add 1 random char from "dashesSlashes" to the variable "tempPWD"
	if useDashesSlashes {
		newBaseString += dashesSlashes
		tempPWD += string(dashesSlashes[rand.Intn(len(dashesSlashes))])
	}

	// Add the charachters from "mathSymbols" to the "newBaseString" variable and Add 1 random char from "mathSymbols" to the variable "tempPWD"
	if useMathSymbols {
		newBaseString += mathSymbols
		tempPWD += string(mathSymbols[rand.Intn(len(mathSymbols))])
	}

	// Add the charachters from "brackets" to the "newBaseString" variable and Add 1 random char from "brackets" to the variable "tempPWD"
	if useBrackets {
		newBaseString += brackets
		tempPWD += string(brackets[rand.Intn(len(brackets))])
	}

	// If the tempPWD exeeds the defiend length, stop the program
	if length <= len(tempPWD) {
		log.Fatal("Password length exceeds the requested length, length must be equal to the amount of flags selected + 1")
		return ""
	}

	// Check how many char are already in tempPWD and add chars from the newBaseString until the length is equal to the "length" defined in flag
	genPWD := make([]byte, length-len(tempPWD))
	for i := range genPWD {
		genPWD[i] = newBaseString[rand.Intn(len(newBaseString))]
	}

	// Combine tempPWD and genPWD (Generated Password) and shuffle the combined string
	newPWD := []byte(tempPWD + string(genPWD))
	for i := range newPWD {
		x := rand.Intn(i + 1)
		newPWD[i], newPWD[x] = newPWD[x], newPWD[i]
	}

	// Return the new password
	return string(newPWD)
}

func PasswordGenerator() string {
	// Variable for the base string that will be used to generate a password
	baseString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Define the length of the password using a flag
	length := flag.Int("L", 16, "How long should the password be?")

	// Define the characters to be included in the password using flags
	useNumbers := flag.Bool("n", false, "Should numbers be used? (Default value: False)")
	useLogograms := flag.Bool("l", false, "Should logograms be used? (Default value: False)")
	usePunctuations := flag.Bool("p", false, "Should punctuation be used? (Default value: False)")
	useQuotations := flag.Bool("q", false, "Should quotations be used? (Default value: False)")
	useDashesSlashes := flag.Bool("ds", false, "Should dashes & slashes be used? (Default value: False)")
	useMathSymbols := flag.Bool("m", false, "Should math symbols be used? (Default value: False)")
	useBrackets := flag.Bool("b", false, "Should brackets be used? (Default value: False)")

	// Parse the above flags
	flag.Parse()

	// Print out the password
	newPassword := Randomize(*length, baseString, *useNumbers, *useLogograms, *usePunctuations, *useQuotations, *useDashesSlashes, *useMathSymbols, *useBrackets)

	return newPassword
}

/*
What need to be done to generate a password that guarantees atleast 1 char of each defined flag is part of the generated password
- Define a basestring
- Define the flags
- Randomise 1 char from each selected flag and add those to the string "newPWD"
- Randomise all the charachters from the basestring + selected flags and add the to the variable "newPWD" till the length of the variable is equal to the defined length
- Randomise the whole "newPWD" again
- Output
*/
