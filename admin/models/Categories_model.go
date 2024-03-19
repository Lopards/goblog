package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&category)

}

func (category Category) Add() { // db ye veri eklemek için

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&category)
	db.Create(&category)
}

func (category Category) Get(where ...interface{}) Category { // db den veri çekmek için  herhangi bir veri gelebilr
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}
	db.First(&category, where...)

	return category
}

func (category Category) GetAll(where ...interface{}) []Category { // tüm verileri çekecek

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)

		return nil
	}
	var categorys []Category
	db.Find(&categorys, where...)
	return categorys
}

func (category Category) Update(colon string, value interface{}) {

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Update(colon, value)

}

func (category Category) Updates(data Category) {

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Updates(data)
}

func (category Category) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&category, category.ID)
}
