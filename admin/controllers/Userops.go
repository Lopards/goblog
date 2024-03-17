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
