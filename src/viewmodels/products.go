package viewmodels

import(
	
)

type Products struct {
	Title string
	Active string
	Products []Product
}

func GetProducts (id int) Products {
	var result Products
	
	result.Active = "shop"
	var shopName string
	
	switch id {
		case 1:
			shopName = "Video Game"
		case 2:
			shopName = "Show"
		case 3:
			shopName = "Video Game"
	}
	
	result.Title = shopName + " shop:"
	
	if id==1 {
		movie1 := MakeMovie("1.jpg", "Ex Machina", `A young programmer is selected to participate in a
						ground-breaking experiment in synthetic intelligence by evaluating
						the human qualities of a breath-taking humanoid A.I.`)
		movie2 := MakeMovie("2.jpg", "Mad Max: Fury Road", `A woman rebels against a tyrannical ruler in postapocalyptic
						Australia in search for her home-land with the help of a group of
						female prisoners, a psychotic worshipper, and a drifter named Max.`)
		
		moviesSlice := []Product{movie1, movie2}
		
		
		result.Products = moviesSlice
	}
	
	return result
}

type Product struct {
	ImageUrl string
	Title string
	Type string
	Description string
}

func MakeMovie(ImageUrl string, Title string, Description string) Product {
	result := Product {ImageUrl, Title, "Movie", Description}
	
	return result
}

func MakeShow(ImageUrl string, Title string, Description string) Product {
	result := Product {ImageUrl, Title, "Show", Description}
	
	return result
}

func MakeVideoGame(ImageUrl string, Title string, Description string) Product {
	result := Product {ImageUrl, Title, "Video Game", Description}
	
	return result
}