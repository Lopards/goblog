package config

import (
	admin "goblog/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//admin
	//Blog Post
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add/", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//kategori
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	// User operasyınları
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)
	r.GET("/admin/register", admin.Userops{}.Register_index)
	r.POST("/admin/do_register", admin.Userops{}.Register)
	// serve files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads")) // blog sayfasındaki kapak görselini sitede göstermek için
	return r
}
