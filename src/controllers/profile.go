package controllers

import (
	"controllers/util"
	"converter"
	"log"
	"models"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetProfile()
	vm.LoggedIn = true

	responseWriter := util.GetResponseWriter(w, req)

	ck, err := req.Cookie("goSessionId")

	if err == nil {
		userId, _ := strconv.Atoi(ck.Value)
		modelMember, _ := models.GetMemberById(userId)
		vm.Member = converter.ConvertMemberlToViewModel(modelMember)

		if req.Method == "GET" {
			var listOfProductIds []int
			listOfProductIds, _ = models.GetMembersOrder(userId)

			var listOfProducts []viewmodels.Product
			for _, val := range listOfProductIds {
				pr, _ := models.GetProduct(val)
				listOfProducts = append(listOfProducts, converter.ConvertProductsToViewModel(pr))
			}

			vm.Products = listOfProducts
			vm.LoggedIn = true

		} else {
			remButton := req.FormValue("remove")
			remId, _ := strconv.Atoi(remButton)

			if remButton == "" {
				productName := req.FormValue("name")
				productType := req.FormValue("type")
				productDescription := req.FormValue("description")
				productPrice := req.FormValue("price")
				productImgUrl := req.FormValue("imageurl")

				_, fileErr := models.GetProductByName(productName)

				if fileErr != nil {
					inputProduct := models.Product{}
					inputProduct.SetName(productName)
					inputProduct.SetDescription(productDescription)
					inputProduct.SetImageUrl(productImgUrl)

					typ, _ := strconv.Atoi(productType)
					inputProduct.SetTyp(typ)

					price64, _ := strconv.ParseFloat(productPrice, 2)
					price := float32(price64)
					inputProduct.SetPrice(price)

					insertErr := models.InsertProduct(inputProduct)

					if insertErr == nil {
						http.Redirect(w, req, "/profile", http.StatusFound)
					} else {
						log.Fatal(insertErr.Error())
					}

				} else {
					http.Redirect(w, req, "/home", http.StatusFound)
				}
			} else {
				deleteErr := models.RemoveOrder(userId, remId)
				
				if deleteErr == nil {
					http.Redirect(w, req, "/profile", http.StatusFound)
				}
			}
		}

	} else {
		vm.LoggedIn = false
	}

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
	defer responseWriter.Close()
}
