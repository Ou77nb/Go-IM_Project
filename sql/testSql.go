package main

import (
	"IM_Project/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/im_project?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Print("gorm.Open:", err)
	}

	// 迁移 schema
	//err2 := db.AutoMigrate(&models.UserBasic{})
	//if err2 != nil {
	//	fmt.Print("db.AutoMigrate:", err2)
	//}
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})
}
