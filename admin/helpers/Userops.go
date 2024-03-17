package helpers

import (
	"fmt"
	"goblog/admin/models"
	"net/http"
)

func Setuser(w http.ResponseWriter, r *http.Request, username string, password string) error {
	session, err := store.Get(r, "blog-user")

	if err != nil {
		fmt.Println(err)
		return err
	}
	session.Values["username"] = username
	session.Values["password"] = password
	return session.Save(r, w)
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		fmt.Println(err)
		return false
	}
	username := session.Values["username"]
	password := session.Values["password"]

	user := models.User{}.Get("username = ? AND password = ?", username, password)
	if user.Username == username && user.Password == password {
		return true
	} else {
		SetAlert(w, r, "Lütfen giriş yapın")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return false
	}
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		fmt.Println(err)
		return err
	}
	session.Options.MaxAge = -1 // burada maxAge ile kullanıcıyı sessiondan dolaylı olarak silmiş olduk (GORİLLA sessionda direkt silme fonksiypnu yok)
	return session.Save(r, w)
}
