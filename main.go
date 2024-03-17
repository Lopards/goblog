package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	//post := admin_models.Post{}.Get(1)
	//post.Updates(admin_models.Post{Title: "go ile web programlama", Description: "test", CategoryID: 23, Content: "23"})
	//post.Update("slug", "web")

	//fmt.Println(admin_models.Post{}.GetAll())

	//fmt.Println(post.Title)
	//post.Delete()
	http.ListenAndServe(":8080", config.Routes())
}
