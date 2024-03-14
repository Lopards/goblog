package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	CategoryID                                     int
}

func (post Post) Migrate() {
	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&post)

}

func (post Post) Add() { // db ye veri eklemek için

	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&post)
	db.Create(&post)
}

func (post Post) Get(where ...interface{}) Post { // db den veri çekmek için  herhangi bir veri gelebilr
	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	db.First(&post, where...)

	return post
}

func (post Post) GetAll(where ...interface{}) []Post { // tüm verileri çekecek

	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)

		return nil
	}
	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (post Post) Update(colon string, value interface{}) {

	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Update(colon, value)

}

func (post Post) Updates(data Post) {

	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete() {
	db, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&post, post.ID)
}
