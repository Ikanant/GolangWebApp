package controllers

import (
	"controllers/util"
	"models"
	"net/http"
	"text/template"
	"viewmodels"
	"converter"
	"io"
	"os"
	"strconv"
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
		if req.Method == "GET" {
			userId, _ := strconv.Atoi(ck.Value)
			modelMember, _ := models.GetMemberById(userId)
			vm.Member = converter.ConvertMemberlToViewModel(modelMember)
			
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
			req.ParseMultipartForm(32 << 20)
			productName := req.FormValue("name")
			productType := req.FormValue("type")
			productDescription := req.FormValue("description")
			productPrice := req.FormValue("price")
			
			productImgFile, _, fileErr := req.FormFile("imageurl")

			if fileErr == nil {
				defer productImgFile.Close()

				println(productName, productType, productDescription, productPrice)

				_, fileErr := models.GetProductByName(productName)

				if fileErr != nil {
					futureId, _ := models.GetNumberOfProducts()
					futureId++

					futureIdStr := strconv.Itoa(futureId) + ".jpg"
					f, _ := os.OpenFile("./public/images/products/"+futureIdStr, os.O_WRONLY|os.O_CREATE, 0666)
					defer f.Close()
					io.Copy(f, productImgFile)
					
					inputProduct := models.Product{}
					inputProduct.SetName(productName)
					inputProduct.SetImageUrl(futureIdStr)
					inputProduct.SetDescription(productDescription)
					
					typ, _ := strconv.Atoi(productType)
					inputProduct.SetTyp(typ)
					
					price64, _ := strconv.ParseFloat(productPrice, 2)
					price := float32(price64)
					inputProduct.SetPrice(price)
					println("=======", price)

					insertErr := models.InsertProduct(inputProduct)

					if insertErr == nil {
						http.Redirect(w, req, "/home", http.StatusFound)
					} else {
						println(insertErr.Error())
					}
					
				} else {
					http.Redirect(w, req, "/home", http.StatusFound)
				}
			} else {
				println(fileErr.Error())
			}
		}

	} else {
		vm.LoggedIn = false
	}

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
	defer responseWriter.Close()
}
