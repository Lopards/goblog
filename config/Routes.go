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
	// serve files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
