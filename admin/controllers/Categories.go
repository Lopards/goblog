package controllers

import (
	"fmt"
	slug2 "github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
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
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (categories Categories) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoyTitle := r.FormValue("category-title")
	CategoySlug := slug2.Make(CategoyTitle)

	models.Category{
		Title: CategoyTitle,
		Slug:  CategoySlug,
	}.Add()
	helpers.SetAlert(w, r, "kayıt başarı ile eklendi")
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}

func (categories Categories) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Silinecek gönderiyi alır ve veritabanından siler.
	post := models.Category{}.Get(params.ByName("id"))
	post.Delete()
	// Kullanıcıyı yönlendirir ve ana sayfaya geri döner.
	helpers.SetAlert(w, r, "kategori  silindi")

	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}
