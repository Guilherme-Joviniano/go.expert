package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey;auto_increment"`
	Name  string
	Price float32
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	connection := mysql.Open(dsn)
	db, err := gorm.Open(connection, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Name: "Monitor AOC HERO", Price: float32(999.90)})

	products := []Product{
		{Name: "Monitor AOC Hero", Price: float32(999.90)},
		{Name: "SpongeBob Toy", Price: float32(49.90)},
		{Name: "Mouse Gamer Hero", Price: float32(129.99)},
	}

	db.Create(&products)


	
}
