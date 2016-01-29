package viewmodels

import(
	
)

type Categories struct {
	Title string
	Active string
	Categories []Category
}

type Category struct {
	Title string
	Id int
	ImageUrl string
	IsOrientedRight bool
}

func GetCategory() Categories {
	result := Categories{
		Title: "Category:",
		Active: "category",
	}
	
	categoriesSlice := []Category{
		Category{"Movies", 1, "frontMovies.png", true},
		Category{"Shows", 2, "frontShows.png", false},
		Category{"Video Games", 3, "frontGames.png", true},
	}
	
	result.Categories = categoriesSlice
	
	return result
}