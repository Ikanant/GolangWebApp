package models

import ()

type Category struct {
	title          string
	id             int
	imageUrl       string
	sampleProducts []string
}

func (this *Category) Title() string {
	return this.title
}
func (this *Category) Id() int {
	return this.id
}
func (this *Category) ImageUrl() string {
	return this.imageUrl
}
func (this *Category) SampleProducts() []string {
	return this.sampleProducts
}
func (this *Category) setTitle(value string) {
	this.title = value
}
func (this *Category) setId(value int) {
	this.id = value
}
func (this *Category) setImageUrl(value string) {
	this.imageUrl = value
}
func (this *Category) setSampleProducts (value []string) {
	this.sampleProducts = value
}

func GetCategories() []Category {
	sampleMovieSlice := []string{"Ex Machina", "Mad Max: Fury Road", "Inside Out", "Star Wars: The Force Awakes", "The Martian"}
	sampleShowSlice := []string{"Daredevil", "Game of Thrones", "Mr. Robot", "Jessica Jones", "Master of None"}
	sampleGameSlice := []string{"The Witcher 3: Wild Hunt", "Metal Gear Solid V", "Bloodborne", "Batman: Arkham Knight", "Fallout 4"}

	categoriesSlice := []Category{
		Category{"Movies", 1, "frontMovies.jpg", sampleMovieSlice},
		Category{"Shows", 2, "frontShows.jpg", sampleShowSlice},
		Category{"Video Games", 3, "frontGames.jpg", sampleGameSlice},
	}
	
	return categoriesSlice 
}
