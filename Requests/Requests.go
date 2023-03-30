package Requests

import (
	"database/sql"
	"log"
	readjson "myapp/Configuration"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (*sql.DB, error) {
	cfg, err := readjson.ReadJson()
	if err != nil {
		log.Fatal(err)
	}

	//DB Login Details'
	dbuser := cfg.DBLogin[0].DBUser
	dbpass := cfg.DBLogin[0].DBPass
	dbname := cfg.DBLogin[0].DBName

	// Make a connection to the DB
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	// If an error occured, return the error
	if err != nil {
		return nil, err
	}
	// Return the db varaible and nil
	return db, nil
}

// Function to check if a password already exists in the database
func CheckPasswordExists(db *sql.DB, pwd string) (bool, error) {
	//Run an GET query to check if the password already exists in the database
	result, err := db.Query("SELECT Password FROM `passwords` WHERE Password = '" + pwd + "'; ")
	if err != nil {
		return false, err
	}

	//If the password exists return true and nil
	if result.Next() {
		return true, nil
	}

	//If the password does not exist return false
	return false, nil
}

// Function to insert a new password into the database
func InsertPassword(db *sql.DB, newPWD string) error {
	//Run an INSERT query to add the new pwd to the database
	insert, err := db.Query("INSERT INTO `passwords` (`id`, `Password`) VALUES (NULL, '" + newPWD + "');")
	if err != nil {
		return err
	}

	//Close the database connection
	defer insert.Close()

	//Return nil
	return nil
}
