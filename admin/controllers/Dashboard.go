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
	if !helpers.CheckUser(w, r) {

		return
	}

	// Gösterilecek sayfanın HTML şablonunu yükler.
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Verileri tutacak bir harita oluşturulur ve bu haritaya tüm gönderiler eklenir.
	data := make(map[string]interface{})
	data["posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	// HTML şablonunu ve verileri kullanarak sayfayı görüntüler.
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	// Yeni öğe ekleme sayfasının HTML şablonunu yükler.
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)

	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()

	// HTML şablonunu görüntüler.
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	// Formdan gelen verileri alır.
	title := r.FormValue("blog-title")
	slug := slug2.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	// Resim dosyasını işler ve sunucuya kaydeder.
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

	// Gönderi modeline yeni bir gönderi ekler.
	models.Post{
		Title:       title,
		Slug:        slug,
		Content:     content,
		Description: description,
		CategoryID:  categoryID,
		Picture_url: "uploads/" + header.Filename,
	}.Add()

	// Kullanıcıyı yönlendirir ve ana sayfaya geri döner.

	helpers.SetAlert(w, r, "Kayıt başarı ile eklendi")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	// Silinecek gönderiyi alır ve veritabanından siler.
	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	// Kullanıcıyı yönlendirir ve ana sayfaya geri döner.
	helpers.SetAlert(w, r, "Kayıt  silindi")

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	// Düzenlenecek gönderinin HTML şablonunu yükler.
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Düzenlenecek gönderiyi alır ve şablonla birlikte görüntüler.
	data := make(map[string]interface{})
	data["post"] = models.Post{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}

	// Güncellenecek gönderiyi veritabanından al.
	post := models.Post{}.Get(params.ByName("id"))

	// Formdan gelen verileri al.
	title := r.FormValue("blog-title")
	slug := slug2.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")

	var picture_url string
	if is_selected == "1" {
		// Yeni bir resim seçilmişse, upload işlemi yapılır.
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
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	} else {
		// Yeni bir resim seçilmemişse, mevcut resmin URL'si kullan.
		picture_url = post.Picture_url
	}

	// Gönderiyi günceller.
	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		CategoryID:  categoryID,
		Content:     content,
		Description: description,
		Picture_url: picture_url,
	})

	helpers.SetAlert(w, r, "Kayıt Güncellendi")
	// Kullanıcıyı güncellenmiş gönderinin düzenleme sayfasına yönlendir.
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
