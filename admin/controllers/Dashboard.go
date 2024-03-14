package controllers

import (
	"fmt"
	slug2 "github.com/gosimple/slug"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"io"
	"net/http"
	"os"
	"strconv"
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

	view.ExecuteTemplate(w, "index", nil)
}

func (dashnoard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title") //formdan gelen verileri almak için
	slug := slug2.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Content:     content,
		Description: description,
		CategoryID:  categoryID,
		Picture_url: "uploads/" + header.Filename,
	}.Add()

	http.Redirect(w, r, "/admin", http.StatusSeeOther) // kayit işlemleri bittikten sonra ana sayfaya donmek için bir komut
	// TODO alert
}
