package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// โครงสร้าง Book ให้ ID เป็น Auto Increment
type Book struct {
	ID     int `gorm:"primaryKey;autoIncrement"` // เพิ่ม autoIncrement
	Title  string
	Author string
	Price  float64
}

func main() {
	dsn := "host=localhost user=postgres password=root1234 dbname=postgres port=9999 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// สร้างตารางใหม่
	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}

	fmt.Println("เชื่อมต่อสำเร็จ!")

	for {
		fmt.Println("\nเลือกการทำงาน:")
		fmt.Println("1. เพิ่มหนังสือ")
		fmt.Println("2. แสดงข้อมูลหนังสือ")
		fmt.Println("3. แก้ไขข้อมูลหนังสือ")
		fmt.Println("4. ลบหนังสือ")
		fmt.Println("5. ออกจากโปรแกรม")

		var choice string
		fmt.Print("เลือกเมนู (1-5): ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			createBook(db)
		case "2":
			readBooks(db)
		case "3":
			updateBook(db)
		case "4":
			deleteBook(db)
		case "5":
			fmt.Println("จบการทำงาน")
			return
		default:
			fmt.Println("กรุณาเลือกเมนู 1-5")
		}
	}
}

// Create
func createBook(db *gorm.DB) {
	var book Book

	fmt.Print("ใส่ชื่อหนังสือ: ")
	fmt.Scanln(&book.Title)

	fmt.Print("ใส่ชื่อผู้แต่ง: ")
	fmt.Scanln(&book.Author)

	fmt.Print("ใส่ราคา: ")
	fmt.Scanln(&book.Price)

	result := db.Create(&book)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	fmt.Printf("เพิ่มหนังสือ ID: %d สำเร็จ\n", book.ID)
}

// Read
func readBooks(db *gorm.DB) {
	fmt.Print("ใส่ ID ที่ต้องการค้นหา (* สำหรับแสดงทั้งหมด): ")
	var input string
	fmt.Scanln(&input)

	var books []Book
	var result *gorm.DB

	if input == "*" {
		result = db.Find(&books)
	} else {
		result = db.Where("id = ?", input).Find(&books)
	}

	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	if len(books) == 0 {
		fmt.Println("ไม่พบข้อมูล")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}

// Update
func updateBook(db *gorm.DB) {
	var id int
	fmt.Print("ใส่ ID ที่ต้องการแก้ไข: ")
	fmt.Scanln(&id)

	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		fmt.Println("ไม่พบหนังสือที่ต้องการแก้ไข")
		return
	}

	fmt.Printf("ข้อมูลเดิม: ID: %d, Title: %s, Author: %s, Price: %.2f\n",
		book.ID, book.Title, book.Author, book.Price)

	fmt.Print("ใส่ชื่อหนังสือใหม่ (Enter เพื่อข้าม): ")
	var title string
	fmt.Scanln(&title)
	if title != "" {
		book.Title = title
	}

	fmt.Print("ใส่ชื่อผู้แต่งใหม่ (Enter เพื่อข้าม): ")
	var author string
	fmt.Scanln(&author)
	if author != "" {
		book.Author = author
	}

	fmt.Print("ใส่ราคาใหม่ (Enter เพื่อข้าม): ")
	var price string
	fmt.Scanln(&price)
	if price != "" {
		fmt.Sscanf(price, "%f", &book.Price)
	}

	result = db.Save(&book)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	fmt.Println("แก้ไขข้อมูลสำเร็จ")
}

// Delete
func deleteBook(db *gorm.DB) {
	var id int
	fmt.Print("ใส่ ID ที่ต้องการลบ: ")
	fmt.Scanln(&id)

	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		fmt.Println("ไม่พบหนังสือที่ต้องการลบ")
		return
	}

	fmt.Printf("ต้องการลบหนังสือ: ID: %d, Title: %s, Author: %s, Price: %.2f\n",
		book.ID, book.Title, book.Author, book.Price)
	fmt.Print("ยืนยันการลบ? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm == "y" {
		result = db.Delete(&book)
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
			return
		}
		fmt.Println("ลบข้อมูลสำเร็จ")
	} else {
		fmt.Println("ยกเลิกการลบ")
	}
}
