package controllers

import (
	"controllers/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
)

type productController struct {
	template         *template.Template
	purchaseTemplate *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	_, er := req.Cookie("goSessionId")

	if er == nil {
		vars := mux.Vars(req)

		idRaw := vars["id"]

		id, err := strconv.Atoi(idRaw)

		if err == nil {
			vm := viewmodels.GetProduct(id)
			vm.LoggedIn = true
			w.Header().Add("Content-Type", "text/html")

			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()
			this.template.Execute(responseWriter, vm)
		} else {
			w.WriteHeader(404)
		}
	} else {
		http.Redirect(w, req, "/home", http.StatusFound)
	}

}
