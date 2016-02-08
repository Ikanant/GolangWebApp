package models

import (
	"errors"
	"log"
	"fmt"
)

type Product struct {
	id          int
	imageUrl    string
	name        string
	typ         int
	description string
	price       float32
}

func (this *Product) Id() int {
	return this.id
}
func (this *Product) SetId(value int) {
	this.id = value
}
func (this *Product) ImageUrl() string {
	return this.imageUrl
}
func (this *Product) SetImageUrl(value string) {
	this.imageUrl = value
}
func (this *Product) Name() string {
	return this.name
}
func (this *Product) SetName(value string) {
	this.name = value
}
func (this *Product) Typ() int {
	return this.typ
}
func (this *Product) SetTyp(value int) {
	this.typ = value
}
func (this *Product) Description() string {
	return this.description
}
func (this *Product) SetDescription(value string) {
	this.description = value
}
func (this *Product) Price() float32 {
	return this.price
}
func (this *Product) SetPrice(value float32) {
	this.price = value
}

func GetProductList(typ int) ([]Product, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		productRows, err2 := db.Query(`SELECT * FROM products WHERE typ = $1;`, typ)

		if err2 == nil {
			productSlice := []Product{}
			var id int
			var imageurl string
			var name string
			var typ int
			var description string
			var price float32

			for productRows.Next() {
				err3 := productRows.Scan(&id, &imageurl, &name, &typ, &description, &price)

				resultP := Product{id, imageurl, name, typ, description, price}

				if err3 == nil {
					productSlice = append(productSlice, resultP)
				} else {
					log.Fatal(err3)
				}
			}

			return productSlice, nil
		} else {
			log.Fatal(err2)
			return nil, errors.New("Couldn't connect to the database")
		}

	} else {
		return nil, errors.New("Couldn't connect to the database")
	}
}

func GetProduct(id int) (Product, error) {
	db, err := getDBConnection()

	product := Product{}
	
	println("++++", id)

	if err == nil {
		defer db.Close()

		productRows := db.QueryRow(`SELECT * FROM products WHERE id = $1;`, id)

		var id int
		var imageurl string
		var name string
		var typ int
		var description string
		var price float32

		err3 := productRows.Scan(&id, &imageurl, &name, &typ, &description, &price)

		fmt.Println("^", name)

		if err3 == nil {
			product.SetId(id)
			product.SetImageUrl(imageurl)
			product.SetName(name)
			product.SetTyp(typ)
			product.SetDescription(description)
			product.SetPrice(price)
		} else {
			log.Fatal(err3)
		}

		return product, nil

	} else {
		return product, errors.New("Couldn't connect to the database")
	}
}

func GetProductByName(name string) (Product, error) {
	db, err := getDBConnection()

	product := Product{}

	if err == nil {
		defer db.Close()

		productRows := db.QueryRow(`SELECT * FROM products WHERE name = $1;`, name)

		if productRows != nil {

			err3 := productRows.Scan(&product.id, &product.imageUrl, &product.name, &product.typ, &product.description, &product.price)
			
			if err3 == nil {
				return product, nil
			} else {
				return product, err3
			}

		}

		return product, nil

	} else {
		return product, errors.New("Couldn't connect to the database")
	}
}

func GetNumberOfProducts() (int, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		productRows := db.QueryRow(`SELECT count(*) FROM products;`)

		var count int

		productRows.Scan(&count)

		return count, nil

	} else {
		return 0, errors.New("Couldn't connect to the database")
	}
}

func InsertProduct(input Product) error {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		sqlObj := db.QueryRow(`INSERT INTO products (id, imageurl, name, typ, description, price) VALUES (DEFAULT, $1, $2, $3, $4, $5);`, input.ImageUrl(), input.Name(), input.Typ(), input.Description(), input.Price())
		
		errI := sqlObj.Scan()
		if errI == nil {
			return nil
		} else {
			println("error:", errI.Error())
			return errI
		}

	} else {
		return errors.New("Couldn't connect to the database")
	}
}
