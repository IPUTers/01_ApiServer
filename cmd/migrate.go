package main

import (
	"try-standard-layout/internal/domain"
	"try-standard-layout/internal/postgres"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		panic(err) // あとで修正
	}
	defer postgres.Close(db)

	// Migrate the schema
	db.AutoMigrate(&domain.TestApi{})

	// 初期データを登録
	db.Create(&domain.TestApi{Name: "Aさん", Age: 19})
	db.Create(&domain.TestApi{Name: "Bさん", Age: 22})
}
