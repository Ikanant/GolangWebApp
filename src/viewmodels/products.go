package viewmodels

import ()

type Products struct {
	Title    string
	Active   string
	Products []Product
}

func GetProducts(id int) Products {
	var result Products

	result.Active = "shop"
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
		movie1 := MakeMovie("movies/1.jpg", "Ex Machina", `A young programmer is selected to participate in a
						ground-breaking experiment in synthetic intelligence by evaluating
						the human qualities of a breath-taking humanoid A.I.`)
		movie2 := MakeMovie("movies/2.jpg", "Mad Max: Fury Road", `A woman rebels against a tyrannical ruler in postapocalyptic
						Australia in search for her home-land with the help of a group of
						female prisoners, a psychotic worshipper, and a drifter named Max.`)
		movie3 := MakeMovie("movies/3.jpg", "Inside Out", `After young Riley is uprooted from her Midwest life and
						moved to San Francisco, her emotions - Joy, Fear, Anger, Disgust
						and Sadness - conflict on how best to navigate a new city, house,
						and school.`)
		movie4 := MakeMovie("movies/4.jpg", "Star Wars: The Force Awakens", `Three decades after the defeat of the Galactic Empire, a new
						threat arises. The First Order attempts to rule the galaxy and
						only a ragtag group of heroes can stop them, along with the help
						of the Resistance.`)
		movie5 := MakeMovie("movies/5.jpg", "The Martian", `During a manned mission to Mars, Astronaut Mark Watney is
						presumed dead after a fierce storm and left behind by his crew.
						But Watney has survived and finds himself stranded and alone on
						the hostile planet. With only meager supplies, he must draw upon
						his ingenuity, wit and spirit to subsist and find a way to signal
						to Earth that he is alive.`)
		movie6 := MakeMovie("movies/6.jpg", "Jurassic World", `A new theme park is built on the original site of Jurassic
						Park. Everything is going well until the parks newest
						attraction--a genetically modified giant stealth killing
						machine--escapes containment and goes on a killing spree.`)

		moviesSlice := []Product{movie1, movie2, movie3, movie4, movie5, movie6}

		result.Products = moviesSlice

	} else if id == 2 {
		show1 := MakeShow("shows/1.jpg", "Daredevil", `A blind lawyer, with his other senses superhumanly enhanced, fights crime as a costumed superhero.`)
		show2 := MakeShow("shows/2.jpg", "Game of Thrones", `While a civil war brews between several noble families in Westeros, the children of the former rulers of the land attempt to rise up to power.`)
		show3 := MakeShow("shows/3.jpg", "Mr. Robot", `Follows a young computer programmer who suffers from social anxiety disorder and forms connections through hacking. He's recruited by a mysterious anarchist, who calls himself Mr. Robot.`)
		show4 := MakeShow("shows/4.jpg", "Jessica Jones", `A former superhero decides to reboot her life by becoming a private investigator.`)
		show5 := MakeShow("shows/5.jpg", "Master of None", `The personal and professional life of Dev, a 30-year-old actor in New York.`)
		
		showsSlice := []Product{show1, show2, show3, show4, show5}
		
		result.Products = showsSlice
	
	} else if id == 3 {
		game1 := MakeVideoGame("games/1.jpg", "Witcher 3: Wild Hunt", `The Witcher 3: Wild Hunt is a high-fantasy, action role-playing video game set in an open-world environment, developed by CD Projekt RED.`)
		game2 := MakeVideoGame("games/2.jpg", "Metal Gear V", `Metal Gear is a series of action-adventure stealth video games, created by Hideo Kojima and developed and published by Konami.`)
		game3 := MakeVideoGame("games/3.jpg", "Bloodborne", `Bloodborne is an action role-playing video game directed by Hidetaka Miyazaki, developed by FromSoftware, and published by Sony Computer Entertainment.`)
		game4 := MakeVideoGame("games/4.jpg", "Batman: Arkham Knight", `Batman: Arkham Knight is a 2015 action-adventure video game developed by Rocksteady Studios and published by Warner Bros. `)
		game5 := MakeVideoGame("games/5.jpg", "Fallout 4", `Fallout 4 is an action role-playing video game developed by Bethesda Game Studios and published by Bethesda Softworks.`)
	
		gamesSlice := []Product{game1, game2, game3, game4, game5}
		
		result.Products = gamesSlice
	}

	return result
}

type Product struct {
	ImageUrl    string
	Title       string
	Type        string
	Description string
}

func MakeMovie(ImageUrl string, Title string, Description string) Product {
	result := Product{ImageUrl, Title, "Movie", Description}

	return result
}

func MakeShow(ImageUrl string, Title string, Description string) Product {
	result := Product{ImageUrl, Title, "Show", Description}

	return result
}

func MakeVideoGame(ImageUrl string, Title string, Description string) Product {
	result := Product{ImageUrl, Title, "Video Game", Description}

	return result
}
