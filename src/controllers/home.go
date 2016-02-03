// This file will hold the back controller functionality for the HOME page

package controllers

import (
	"controllers/util"
	"models"
	"net/http" // Will allow us to work with the response writer and request object
	"strconv"
	"text/template" // Will enable us to work with the VIEWs templates
	"time"
	"viewmodels" // Get the View model that provide the HOME page with its data
)

/*
Notice that both the HOME controller and its field begin with LOWER case letters
--> This means that they are scoped to the controller package. This will keep the public API of the controller package
	nice and light since nothing else really needs to know how the controller layer goes about organzing itself
*/
type homeController struct {
	template       *template.Template
	loginTemplate  *template.Template
	signupTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	// Since this controller method is only responsible for a single VIEW, its implementation becomes really simple

	// Grab a view model that is preconfigured with MOCK data using the VIEW model GETHOME
	vm := viewmodels.GetHome()

	// Tell the browser that we will be sending HTML by setting the content-type header on the response
	w.Header().Add("Content-Type", "text/html")

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	_, err := req.Cookie("goSessionId")
	if err == nil {
		vm.LoggedIn = true
		
		ckn, _ := req.Cookie("loggedName")
		
		vm.LoggedName = ckn.Value
	} else {
		vm.LoggedIn = false
	}

	/* We execute the controllers template field.
	This small/lite controlled funciton is important
	Due the controllers position in the MVC pattern it is aware of both of the other layer. This makes it easy to use the
	controller to handle things that really are the responsability of another portion of the application */
	this.template.Execute(responseWriter, vm)

}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")

	vm := viewmodels.GetLogin()

	if req.FormValue("submit") == "signup" {
		http.Redirect(w, req, "/signup", http.StatusFound)
	} else {

		if req.Method == "POST" {

			email := req.FormValue("email")
			password := req.FormValue("password")

			member, err := models.GetMember(email, password)

			if err == nil {
				session, err := models.CreateSession(member)

				if err == nil {
					var cookie http.Cookie
					cookie.Name = "goSessionId"
					cookie.Expires = time.Now().Add(10 * time.Minute)
					cookie.Value = strconv.Itoa(session.MemberId())
					responseWriter.Header().Add("Set-Cookie", cookie.String())
					
					var cookie2 http.Cookie
					cookie2.Name = "loggedName"
					cookie2.Expires = time.Now().Add(10 * time.Minute)
					cookie2.Value = member.FirstName()
					responseWriter.Header().Add("Set-Cookie", cookie2.String())
				}
				vmh := viewmodels.GetHome()

				vmh.LoggedIn = true

				this.template.Execute(responseWriter, vmh)
			} else {
				this.loginTemplate.Execute(responseWriter, vm)
			}
		} else {
			this.loginTemplate.Execute(responseWriter, vm)
		}
	}
}

func (this *homeController) signup(w http.ResponseWriter, req *http.Request) {

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")

	if req.Method == "POST" {
		firstName := req.FormValue("firstName")
		email := req.FormValue("email")
		password := req.FormValue("password")

		err := models.InsertMember(firstName, email, password)

		if err == nil {
			member, _ := models.GetMember(email, password)

			member, errr := models.GetMember(email, password)

			if errr == nil {
				session, err := models.CreateSession(member)

				if err == nil {
					var cookie http.Cookie
					cookie.Name = "goSessionId"
					cookie.Expires = time.Now().Add(10 * time.Minute)
					cookie.Value = strconv.Itoa(session.MemberId())
					responseWriter.Header().Add("Set-Cookie", cookie.String())
					
					var cookie2 http.Cookie
					cookie2.Name = "loggedName"
					cookie2.Expires = time.Now().Add(10 * time.Minute)
					cookie2.Value = member.FirstName()
					responseWriter.Header().Add("Set-Cookie", cookie2.String())
				}

				http.Redirect(w, req, "/home", http.StatusFound)
			}
		}
	}

	vm := viewmodels.GetSignup()

	this.signupTemplate.Execute(responseWriter, vm)

}

func (this *homeController) logout(w http.ResponseWriter, req *http.Request) {

	ck, err := req.Cookie("goSessionId")

	if err == nil {
		removedS := models.RemoveSession(ck.Value)
		if removedS {
			cookieMonster := &http.Cookie{
				Name:    "goSessionId",
				Expires: time.Now(),
				Value:   strconv.FormatInt(time.Now().Unix(), 10),
			}
			cookieMonster2 := &http.Cookie{
				Name:    "loggedName",
				Expires: time.Now(),
				Value:   strconv.FormatInt(time.Now().Unix(), 10),
			}
			
			http.SetCookie(w, cookieMonster)
			http.SetCookie(w, cookieMonster2)

			http.Redirect(w, req, "/home", http.StatusFound)
		}
	}
}
