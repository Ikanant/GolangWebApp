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

/*
package models

// The undescore is important since we aren't actually referencing the package directly in our code.
// Rather we are importing the package for side effects only
// We want to use the library inicialization code, but we don't need anything esle
import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func getDBConnection() (*sql.DB, error) {

	dbUrl := "postgres://smjmqpgjbdlofw:Cq0Pxs_LRU0x1XDHq4Fv1IXLBe@ec2-107-21-218-93.compute-1.amazonaws.com:5432/d9qkjrujasg2kn"
	connection, _ := pq.ParseURL(dbUrl)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Println(err)
	} 

	return db, err
}
*/