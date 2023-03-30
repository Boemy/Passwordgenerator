package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	Requests "myapp/Requests"
)

// Variable for the base string that will be used to generate a password
var baseString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function to randomize the password string
func Randomize(length int, baseString string) string {
	char := make([]byte, length)
	for i := range char {
		char[i] = baseString[rand.Intn(len(baseString))]
	}
	return string(char)
}

func main() {
	// Define the length of the password using a flag
	length := flag.Int("L", 16, "How long should the password be? (Default value: 16)")

	// Define the characters to be included in the password using flags
	numbers := flag.Bool("n", false, "Should numbers be used? (Default value: False)")
	logograms := flag.Bool("l", false, "Should logograms be used? (Default value: False)")
	punctuations := flag.Bool("p", false, "Should punctuation be used? (Default value: False)")
	quotations := flag.Bool("q", false, "Should quotations be used? (Default value: False)")
	dashesSlashes := flag.Bool("ds", false, "Should dashes & slashes be used? (Default value: False)")
	mathSymbols := flag.Bool("m", false, "Should math symbols be used? (Default value: False)")
	brackets := flag.Bool("b", false, "Should brackets be used? (Default value: False)")

	// Parse the above flags
	flag.Parse()

	// Add characters to the base string based on the flag values
	if *numbers {
		baseString += "1234567890"
	}
	if *logograms {
		baseString += "#$%&@^~`"
	}
	if *punctuations {
		baseString += ".,:;"
	}
	if *quotations {
		baseString += `"'`
	}
	if *dashesSlashes {
		baseString += `\/|_-`
	}
	if *mathSymbols {
		baseString += "<>*+!?="
	}
	if *brackets {
		baseString += "{[()]}"
	}

	// Print out the password
	NewPassword := Randomize(*length, baseString)

	// Check if the password already exists in the database
	db, err := Requests.DBConn()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}

	//Close the database connection after it has been used
	defer db.Close()

	//Check if the password exists in the database
	CheckPWD, err := Requests.CheckPasswordExists(db, NewPassword)
	if err != nil {
		log.Println("Error checking if password exists in database:", err)
		return
	}

	//If CheckPWD returned false or error, add password to database
	if !CheckPWD {
		err = Requests.InsertPassword(db, NewPassword)
		if err != nil {
			log.Println("Error inserting password into database:", err)
			return
		}
		// Print a msg with the new password
		fmt.Println("The following password has been added to the database: " + NewPassword)
	} else {
		fmt.Println("Password already exists in database and was not added, your password was: " + NewPassword)
	}

}
