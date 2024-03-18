package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"net/http"
	"text/template"
)

type Userops struct {
}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)

}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password")))) //şifreyi hashledik

	user := models.User{}.Get("username = ? And password = ?", username, password)
	fmt.Println(user)
	if user.Username == username && user.Password == password {
		//login yapıldı
		helpers.Setuser(w, r, username, password)
		helpers.SetAlert(w, r, "Hoş Geldiniz")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {
		helpers.SetAlert(w, r, "yanlış kullanici adı veya şifre girildi")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

}

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	helpers.Logout(w, r)
	helpers.SetAlert(w, r, "Çıkış yaptınız")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func (userops Userops) Register_index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/register")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "register", data)
}
func (userops Userops) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("username = ?", username)

	if user.Username == username {
		fmt.Println("Bu kullanıcı adı kullanılmakta")
		helpers.SetAlert(w, r, "Bu kullanıcı adı kullanılmakta. Lütfen başka bir kullanıcı adı deneyiniz.")
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}

	models.User{
		Username: username,
		Password: password,
	}.Add()

	// Kayıt tamamlandıktan sonra boş bir index sayfası gösterilir
	dashboard := Dashboard{}
	dashboard.Index(w, r, params)
}
