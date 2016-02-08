package controllers

import (
	"controllers/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
	"models"
	"converter"
	"fmt"
)

type productController struct {
	template         *template.Template
	purchaseTemplate *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	ck, er := req.Cookie("goSessionId")

	if er == nil {
		vars := mux.Vars(req)
		idRaw := vars["id"]
		id, err := strconv.Atoi(idRaw)

		if req.Method == "POST" {
			userId, _ := strconv.Atoi(ck.Value)
			productId := id
			
			models.InsertOrder(userId, productId)
			
			http.Redirect(w, req, "/categories", http.StatusFound)
		} else {
			if err == nil {
				vm := viewmodels.GetProductVM()
				vm.LoggedIn = true
				
				modelProduct, _ := models.GetProduct(id)
				vm.Product = converter.ConvertProductsToViewModel(modelProduct)
				
				w.Header().Add("Content-Type", "text/html")

				responseWriter := util.GetResponseWriter(w, req)
				defer responseWriter.Close()
				
				fmt.Println(vm)
				
				this.template.Execute(responseWriter, vm)
			} else {
				w.WriteHeader(404)
			}
		}
	} else {
		http.Redirect(w, req, "/home", http.StatusFound)
	}

}
