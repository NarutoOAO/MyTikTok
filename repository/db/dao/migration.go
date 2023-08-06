package dao

import (
	"MyTikTok/repository/db/model"
	"fmt"
	"os"
)

// Migration Implementation of data migration
func Migration() {
	//Implementation of data migration
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Video{}, &model.User{}, &model.Comment{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
