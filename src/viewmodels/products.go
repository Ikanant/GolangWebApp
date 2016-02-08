package viewmodels

import ()

type Products struct {
	Title      string
	Active     string
	Products   []Product
	LoggedIn   bool
	LoggedName string
}

func GetProducts() Products {
	result := Products{
		Title:    "Products",
		Active:   "category",
		LoggedIn: false,
	}
	return result
}

type ProductVM struct {
	Title      string
	Active     string
	Product    Product
	LoggedIn   bool
	LoggedName string
}

type Product struct {
	Id          int
	ImageUrl    string
	Name        string
	Type        int
	Description string
	Price       float32
}

func GetProductVM() ProductVM {
	result := ProductVM{
		Title:    "Product",
		Active:   "category",
		LoggedIn: false,
	}
	return result
}
