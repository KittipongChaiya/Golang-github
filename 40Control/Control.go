package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Binding from JSON
type Register struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

// สำหรับข้อมูลที่จะส่งกลับไปที่ client (ไม่มี password)
type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
}

// แปลง Register เป็น RegisterResponse
func (r *Register) ToResponse() RegisterResponse {
	return RegisterResponse{
		ID:       r.ID,
		Username: r.Username,
		Fullname: r.Fullname,
		Avatar:   r.Avatar,
	}
}

// HashPassword เข้ารหัสรหัสผ่านด้วย bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword ตรวจสอบรหัสผ่าน
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	dsn := "host=localhost user=postgres password=root1234 dbname=postgres port=9999 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// สร้างตารางใหม่
	err = db.AutoMigrate(&Register{})
	if err != nil {
		panic(err)
	}

	fmt.Println("เชื่อมต่อสำเร็จ!")

	for {
		fmt.Println("\nเลือกการทำงาน:")
		fmt.Println("1. เพิ่ม Username")
		fmt.Println("2. แสดงข้อมูลทั้งหมด")
		fmt.Println("3. ตรวจสอบรหัสผ่าน")
		//fmt.Println("4. ลบหนังสือ")
		fmt.Println("5. ออกจากโปรแกรม")

		var choice string
		fmt.Print("เลือกเมนู (1-5): ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			createregister(db)
		case "2":
			readregisters(db)
		case "3":
			verifyPassword(db)
		case "5":
			fmt.Println("จบการทำงาน")
			return
		default:
			fmt.Println("กรุณาเลือกเมนู 1-5")
		}
	}
}

// Create
func createregister(db *gorm.DB) {
	var register Register

	fmt.Print("ใส่ชื่อUsername: ")
	fmt.Scanln(&register.Username)

	fmt.Print("ใส่รหัสผ่าน: ")
	fmt.Scanln(&register.Password)

	// เข้ารหัสรหัสผ่านก่อนบันทึก
	hashedPassword, err := HashPassword(register.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	register.Password = hashedPassword

	fmt.Print("ใส่ชื่อ-นามสกุล: ")
	fmt.Scanln(&register.Fullname)

	fmt.Print("ใส่ที่อยู่รูปโปรไฟล์: ")
	fmt.Scanln(&register.Avatar)

	result := db.Create(&register)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	fmt.Printf("เพิ่มหนังสือ ID: %d สำเร็จ\n", register.ID)
}

// Read
func readregisters(db *gorm.DB) {
	fmt.Print("ใส่ ID ที่ต้องการค้นหา (* สำหรับแสดงทั้งหมด): ")
	var input string
	fmt.Scanln(&input)

	var registers []Register
	var result *gorm.DB

	if input == "*" {
		result = db.Find(&registers)
	} else {
		result = db.Where("id = ?", input).Find(&registers)
	}

	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	if len(registers) == 0 {
		fmt.Println("ไม่พบข้อมูล")
		return
	}

	for _, register := range registers {
		fmt.Printf("ID: %d, Username: %s, Password: %s, Fullname: %s,Avatar: %s\n",
			register.ID, register.Username, register.Password, register.Fullname, register.Avatar)
	}

}

// VerifyPassword ฟังก์ชันสำหรับตรวจสอบรหัสผ่าน
func verifyPassword(db *gorm.DB) {
	var username string
	var password string
	var register Register

	fmt.Print("ใส่ Username: ")
	fmt.Scanln(&username)

	fmt.Print("ใส่รหัสผ่านที่ต้องการตรวจสอบ: ")
	fmt.Scanln(&password)

	// ค้นหาผู้ใช้จากฐานข้อมูล
	result := db.Where("username = ?", username).First(&register)
	if result.Error != nil {
		fmt.Println("ไม่พบผู้ใช้")
		return
	}

	// ตรวจสอบรหัสผ่าน
	if CheckPassword(password, register.Password) {
		fmt.Println("รหัสผ่านถูกต้อง")
	} else {
		fmt.Println("รหัสผ่านไม่ถูกต้อง")
	}
}
