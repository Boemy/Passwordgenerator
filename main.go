package main

import (
	"fmt"
	"log"
	passwordgenerator "myapp/Passwordgenerator"
	Requests "myapp/Requests"
)

func main() {
	// Check if the password already exists in the database
	db, err := Requests.DBConn()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}

	//Close the database connection after it has been used
	defer db.Close()

	// Genearte a new password
	newPassword := passwordgenerator.PasswordGenerator()

	//Check if the password exists in the database
	CheckPWD, err := Requests.CheckPasswordExists(db, newPassword)
	if err != nil {
		log.Println("Error checking if password exists in database:", err)
		return
	}

	//If CheckPWD returned false or error, add password to database
	if !CheckPWD {
		err = Requests.InsertPassword(db, newPassword)
		if err != nil {
			log.Println("Error inserting password into database:", err)
			return
		}
		// Print a msg with the new password
		fmt.Println("The following password has been added to the database: " + newPassword)
	} else {
		fmt.Println("Password already exists in database and was not added, your password was: " + newPassword)
	}

}
