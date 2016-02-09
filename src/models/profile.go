package models

import (
	"errors"
	"log"
)

func InsertOrder(userId int, productId int) error {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		db.QueryRow(`INSERT INTO orderHistory(userId, productId) VALUES ($1, $2);`, userId, productId)

		return nil
	} else {
		return errors.New("Couldn't connect to the database")
	}
}

func GetMembersOrder(userId int) ([]int, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		orderRows, err2 := db.Query(`SELECT * FROM orderhistory WHERE userid = $1;`, userId)

		if err2 == nil {
			orderIdSlice := []int{}
			var userId int
			var productId int
			for orderRows.Next() {
				err3 := orderRows.Scan(&userId, &productId)
				if err3 == nil {
					orderIdSlice = append(orderIdSlice, productId)
				} else {
					log.Fatal(err3)
				}
			}
			
			return orderIdSlice, nil
		} else {
			log.Fatal(err2)
			return nil, errors.New("Couldn't connect to the database")
		}

	} else {
		return nil, errors.New("Couldn't connect to the database")
	}
}

func RemoveOrder (userId int, productId int) error {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		db.QueryRow(`DELETE FROM orderHistory WHERE userid = $1 AND productid = $2;`, userId, productId);
			
		return nil
		
	} else {
		return errors.New("Couldn't connect to the database")
	}
}
