package viewmodels

import ()

type Profile struct {
	Title    string
	Active   string
	Products []Product
	Member   Member
	LoggedIn bool
}

type Member struct {
	Email     string
	Id        int
	FirstName string
}

func GetProfile() Profile {
	result := Profile{
		Title:  "Profile",
		Active: "profile",
	}

	return result
}
