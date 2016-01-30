package viewmodels

import (

)

type Profile struct {
	Title string
	Active string
	User User
}

type User struct {
	Id int
	Email string
	First string
	Last string
}

func GetProfile() Profile {
	result := Profile {
		Title: "Profile",
	}
	
	return result
}