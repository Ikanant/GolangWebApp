package models

// The undescore is important since we aren't actually referencing the package directly in our code.
// Rather we are importing the package for side effects only
// We want to use the library inicialization code, but we don't need anything esle
import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func getDBConnection() (*sql.DB, error) {
	
	DB_USER := "postgres"
	DB_PASSWORD := "apple"
	DB_NAME := "GotrainDB"
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}
