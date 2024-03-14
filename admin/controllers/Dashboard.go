package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (dashnoard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)

	if err != nil {
		fmt.Println(err)
	}

	view.ExecuteTemplate(w, "index",nil)
}
