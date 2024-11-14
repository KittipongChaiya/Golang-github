package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func main() {
	// เชื่อมต่อกับ PostgreSQL
	connStr := "host=localhost port=9999 user=postgres password=root1234 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ทดสอบการเชื่อมต่อ
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("เชื่อมต่อสำเร็จ!")

	fmt.Print("ใส่ ID ที่ต้องการค้นหา (* สำหรับแสดงทั้งหมด): ")
	var input string
	fmt.Scanln(&input)

	var query string
	if input == "*" {
		query = "SELECT id, title, author, price FROM Book"
	} else {
		query = fmt.Sprintf("SELECT id, title, author, price FROM Book WHERE id = %s", input)
	}

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
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
