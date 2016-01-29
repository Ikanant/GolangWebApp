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
	
	sampleMovieSlice := []string{"Ex Machina", "Mad Max: Fury Road", "Inside Out", "Star Wars: The Force Awakes", "The Martian"}
	sampleShowSlice := []string{"Daredevil", "Game of Thrones", "Mr. Robot", "Jessica Jones", "Master of None"}
	sampleGameSlice := []string{"The Witcher 3: Wild Hunt", "Metal Gear Solid V", "Bloodborne", "Batman: Arkham Knight", "Fallout 4"}

	categoriesSlice := []Category{
		Category{"Movies", 1, "frontMovies.jpg", sampleMovieSlice},
		Category{"Shows", 2, "frontShows.jpg", sampleShowSlice},
		Category{"Video Games", 3, "frontGames.jpg", sampleGameSlice},
	}

	result.Categories = categoriesSlice

	return result
}
