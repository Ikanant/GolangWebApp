package viewmodels

import ()

type Categories struct {
	Title      string
	Active     string
	Categories []Category
	LoggedIn bool
	LoggedName string
}

type Category struct {
	Title           string
	Id              int
	ImageUrl        string
	SampleProducts  []string
}

func GetCategory() Categories {
	result := Categories{
		Title:  "Categories page",
		Active: "category",
	}
	
	return result
}
