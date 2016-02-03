package viewmodels

import ()

type Products struct {
	Title    string
	Active   string
	Products []Product
	LoggedIn bool
	LoggedName string
}

func GetProducts(id int) Products {
	var result Products

	result.Active = "category"
	var shopName string

	switch id {
	case 1:
		shopName = "Movies"
	case 2:
		shopName = "Shows"
	case 3:
		shopName = "Video Games"
	}

	result.Title = shopName

	if id == 1 {
		result.Products = getMoviesProductList()

	} else if id == 2 {
		result.Products = getShowsProductList()

	} else if id == 3 {
		result.Products = getGamesProductList()
	}

	result.LoggedIn = false
	return result
}

func getMoviesProductList() []Product {
	movie1 := MakeMovie(1001, "movies/1.jpg", "Ex Machina", `A young programmer is selected to participate in a
						ground-breaking experiment in synthetic intelligence by evaluating
						the human qualities of a breath-taking humanoid A.I.`)
	movie2 := MakeMovie(1002, "movies/2.jpg", "Mad Max: Fury Road", `A woman rebels against a tyrannical ruler in postapocalyptic
						Australia in search for her home-land with the help of a group of
						female prisoners, a psychotic worshipper, and a drifter named Max.`)
	movie3 := MakeMovie(1003, "movies/3.jpg", "Inside Out", `After young Riley is uprooted from her Midwest life and
						moved to San Francisco, her emotions - Joy, Fear, Anger, Disgust
						and Sadness - conflict on how best to navigate a new city, house,
						and school.`)
	movie4 := MakeMovie(1004, "movies/4.jpg", "Star Wars: The Force Awakens", `Three decades after the defeat of the Galactic Empire, a new
						threat arises. The First Order attempts to rule the galaxy and
						only a ragtag group of heroes can stop them, along with the help
						of the Resistance.`)
	movie5 := MakeMovie(1005, "movies/5.jpg", "The Martian", `During a manned mission to Mars, Astronaut Mark Watney is
						presumed dead after a fierce storm and left behind by his crew.
						But Watney has survived and finds himself stranded and alone on
						the hostile planet. With only meager supplies, he must draw upon
						his ingenuity, wit and spirit to subsist and find a way to signal
						to Earth that he is alive.`)
	movie6 := MakeMovie(1006, "movies/6.jpg", "Jurassic World", `A new theme park is built on the original site of Jurassic
						Park. Everything is going well until the parks newest
						attraction--a genetically modified giant stealth killing
						machine--escapes containment and goes on a killing spree.`)

	moviesSlice := []Product{movie1, movie2, movie3, movie4, movie5, movie6}

	return moviesSlice
}

func getShowsProductList() []Product {
	show1 := MakeShow(2001, "shows/1.jpg", "Daredevil", `A blind lawyer, with his other senses superhumanly enhanced, fights crime as a costumed superhero.`)
	show2 := MakeShow(2002, "shows/2.jpg", "Game of Thrones", `While a civil war brews between several noble families in Westeros, the children of the former rulers of the land attempt to rise up to power.`)
	show3 := MakeShow(2003, "shows/3.jpg", "Mr. Robot", `Follows a young computer programmer who suffers from social anxiety disorder and forms connections through hacking. He's recruited by a mysterious anarchist, who calls himself Mr. Robot.`)
	show4 := MakeShow(2004, "shows/4.jpg", "Jessica Jones", `A former superhero decides to reboot her life by becoming a private investigator.`)
	show5 := MakeShow(2005, "shows/5.jpg", "Master of None", `The personal and professional life of Dev, a 30-year-old actor in New York.`)

	showsSlice := []Product{show1, show2, show3, show4, show5}

	return showsSlice
}

func getGamesProductList() []Product {
	game1 := MakeVideoGame(3001, "games/1.jpg", "Witcher 3: Wild Hunt", `The Witcher 3: Wild Hunt is a high-fantasy, action role-playing video game set in an open-world environment, developed by CD Projekt RED.`)
	game2 := MakeVideoGame(3002, "games/2.jpg", "Metal Gear V", `Metal Gear is a series of action-adventure stealth video games, created by Hideo Kojima and developed and published by Konami.`)
	game3 := MakeVideoGame(3003, "games/3.jpg", "Bloodborne", `Bloodborne is an action role-playing video game directed by Hidetaka Miyazaki, developed by FromSoftware, and published by Sony Computer Entertainment.`)
	game4 := MakeVideoGame(3004, "games/4.jpg", "Batman: Arkham Knight", `Batman: Arkham Knight is a 2015 action-adventure video game developed by Rocksteady Studios and published by Warner Bros. `)
	game5 := MakeVideoGame(3005, "games/5.jpg", "Fallout 4", `Fallout 4 is an action role-playing video game developed by Bethesda Game Studios and published by Bethesda Softworks.`)

	gamesSlice := []Product{game1, game2, game3, game4, game5}

	return gamesSlice
}

type ProductVM struct {
	Title   string
	Active  string
	Type	int
	Product Product
	LoggedIn bool
}

func GetProduct(id int) ProductVM {
	var result ProductVM
	
	var productList []Product
	var product Product
	var typ int
	
	if id>=1000 && id<2000 {
		productList = getMoviesProductList()
		typ = 1
	} else if id >=2000 && id <3000 {
		productList = getShowsProductList()
		typ = 2
	} else if id >=3000 && id <4000 {
		productList = getGamesProductList()
		typ = 3
	}

	for _, p := range productList {
		if p.Id == id {
			product = p
			break
		}
	}

	result.Active = "category"
	result.Title = "Purchase: " + product.Name
	result.Type = typ
	result.Product = product
	result.LoggedIn = false
	
	return result
}

type Product struct {
	Id          int
	ImageUrl    string
	Name        string
	Type        string
	Description string
	Price		float32
}

func MakeMovie(Id int, ImageUrl string, Name string, Description string) Product {
	result := Product{Id, ImageUrl, Name, "Movie", Description, 19.99}

	return result
}

func MakeShow(Id int, ImageUrl string, Name string, Description string) Product {
	result := Product{Id, ImageUrl, Name, "Show", Description, 39.99}

	return result
}

func MakeVideoGame(Id int, ImageUrl string, Name string, Description string) Product {
	result := Product{Id, ImageUrl, Name, "Video Game", Description, 59.99}

	return result
}
