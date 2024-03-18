package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"net/http"
	"text/template"
)

type Categories struct {
}

func (categories Categories) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("Categories/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
