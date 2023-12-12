package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float32
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	connection := mysql.Open(dsn)
	db, err := gorm.Open(connection, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	if err != nil {
		panic(err)
	}

	// db.Create(&Product{Name: "Monitor AOC HERO", Price: float32(999.90)})

	// products := []Product{
	// 	{Name: "Monitor AOC Hero", Price: float32(999.90)},
	// 	{Name: "SpongeBob Toy", Price: float32(49.90)},
	// 	{Name: "Mouse Gamer Hero", Price: float32(129.99)},
	// }

	// db.Create(&products)

	// Select's
	// var product Product
	// db.First(&product, 1)
	// log.Println(product)

	// var products []Product
	// db.Find(&products)
	// log.Println(products)

	// select with where
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// log.Println(products)

	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// log.Println(products)

	// db.Where("id = ?", 1).Find(&products)
	// log.Println(products)

	// Edit and Delete

	// var product Product
	// db.First(&product, 1)
	// product.Name = "New mouse"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 1)
	// log.Println(product2)

	// db.Delete(&product2)

	// Soft Delete
	// product := Product{
	// 	Name:  "Monitor AOC Hero",
	// 	Price: float32(1000),

	// }
	// db.Create(&product)
	// db.Delete(&product)

	// Belongs To

	//category := Category{
	//	Name: "Eletronics",
	//}
	//
	//db.Create(&category)
	//
	//db.Create(&Product{
	//	Name:       "Macbook M1 Air 2020",
	//	Price:      float32(10000),
	//	CategoryID: category.ID,
	//})
	//
	//var products []Product
	//db.Find(&products)
	//
	// for _, product := range products {
	//	log.Println(product)
	//}

	//  Has one

	//category := Category{
	//	Name: "Eletronics",
	//}
	//
	//db.Create(&category)
	//
	//db.Create(&Product{
	//	Name:       "Macbook M1 Air 2020",
	//	Price:      float32(10000),
	//	CategoryID: category.ID,
	//})
	//
	//db.Create(&SerialNumber{
	//	Number:    "a9080938210931283",
	//	ProductId: 1,
	//})
	//
	//var products []Product
	//
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}

	// has many
	//var categories []Category
	//err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.Name)
	//	for _, product := range category.Products {
	//		fmt.Println("- ", product.Name)
	//	}
	//}

	// many to many

	//category := Category{Name: "kitchen"}
	//category2 := Category{Name: "eletronic"}
	//
	//db.Create(&category)
	//
	//db.Create(&Product{
	//	Name:  "Fogão",
	//	Price: float32(3200.99),
	//	SerialNumber: SerialNumber{
	//		Number: "123",
	//	},
	//	Categories: []Category{
	//		category,
	//		category2,
	//	},
	//})
	//
	//var categories []Category
	//
	//err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.Name)
	//	for _, product := range category.Products {
	//		fmt.Println("- ", product.Name)
	//	}
	//}

	// LOCKS

}
