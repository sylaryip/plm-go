package main

import "go-plm/libs/db"

func main() {
	db.InitDB()

	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
