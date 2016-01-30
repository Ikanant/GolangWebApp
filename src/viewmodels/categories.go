package viewmodels

import ()

type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	Title           string
	Id              int
	ImageUrl        string
	SampleProducts  []string
}

func GetCategory() Categories {
	result := Categories{
		Title:  "Category:",
		Active: "category",
	}
	
	return result
}
