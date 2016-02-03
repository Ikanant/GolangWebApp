package viewmodels

import()

type Home struct {
	Title string
	Active string
	LoggedIn bool
	LoggedName string
}

func GetHome() Home {
	result := Home{
		Title: "Home page",
		Active: "home",
	}
	result.LoggedIn = false
	
	return result
}

type Login struct {
	Title string
	Active string
	LoggedIn bool
}

func GetLogin() Login{
	result := Login {
		Title: "Login page",
		Active: "login",
	}
	
	return result
}

type Signup struct {
	Title string
	Active string
	LoggedIn bool
}

func GetSignup() Signup{
	result := Signup {
		Title: "Signup page",
		Active: "signup",
	}
	result.LoggedIn = false
	
	return result
}