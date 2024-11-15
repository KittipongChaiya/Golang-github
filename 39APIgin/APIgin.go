package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	// POST /register - สร้าง user ใหม่
	r.POST("/register", func(c *gin.Context) {
		var register Register
		if err := c.ShouldBindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ตรวจสอบ username ซ้ำ
		var existingUser Register
		result := db.Where("username = ?", register.Username).First(&existingUser)
		if result.Error == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username นี้มีอยู่แล้ว",
			})
			return
		}

		// เข้ารหัส password ด้วย bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเข้ารหัสผ่านได้"})
			return
		}
		register.Password = string(hashedPassword)

		// บันทึกลงฐานข้อมูล
		result = db.Create(&register)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// ส่งข้อมูลกลับโดยไม่รวม password
		c.JSON(http.StatusOK, gin.H{
			"message": "สมัครสมาชิกสำเร็จ!",
			"data":    register.ToResponse(),
		})
	})

	// GET /register - แสดงข้อมูลทั้งหมด (ไม่รวม password)
	r.GET("/register", func(c *gin.Context) {
		var registers []Register
		result := db.Find(&registers)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// แปลงข้อมูลให้ไม่มี password ก่อนส่งกลับ
		var responses []RegisterResponse
		for _, reg := range registers {
			responses = append(responses, reg.ToResponse())
		}

		c.JSON(http.StatusOK, gin.H{
			"data": responses,
		})
	})

	// GET /register/:id - แสดงข้อมูลตาม ID (ไม่รวม password)
	r.GET("/register/:id", func(c *gin.Context) {
		id := c.Param("id")
		var register Register

		result := db.First(&register, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ไม่เจอข้อมูล!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": register.ToResponse(),
		})
	})

	// GET /check-username/:username - เช็ค username ว่ามีอยู่แล้วหรือไม่
	r.GET("/check-username/:username", func(c *gin.Context) {
		username := c.Param("username")
		var register Register

		result := db.Where("username = ?", username).First(&register)
		exists := result.Error == nil

		c.JSON(http.StatusOK, gin.H{
			"มี Username นี้แล้ว": exists,
		})
	})

	// POST /login - สำหรับ login
	r.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ค้นหา user จาก username
		var user Register
		result := db.Where("username = ?", loginData.Username).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username หรือ Password ไม่ตรงกัน"})
			return
		}

		// เปรียบเทียบ password
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username หรือ Password ไม่ตรงกัน"})
			return
		}

		// login สำเร็จ
		c.JSON(http.StatusOK, gin.H{
			"message": "Login สำเร็จ",
			"data":    user.ToResponse(),
		})
	})

	// PUT /register/:id - อัพเดทข้อมูล
	r.PUT("/register/:id", func(c *gin.Context) {
		id := c.Param("id")

		// ตรวจสอบว่ามีข้อมูลอยู่หรือไม่
		var register Register
		if err := db.First(&register, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล!"})
			return
		}

		// รับข้อมูลใหม่จาก request body
		var updateData Register
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// อัพเดทข้อมูล
		result := db.Model(&register).Updates(map[string]interface{}{
			"username": updateData.Username,
			"password": updateData.Password,
			"fullname": updateData.Fullname,
			"avatar":   updateData.Avatar,
		})

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Updated ข้อมูลสำเร็จแล้ว",
			"data":    register,
		})
	})

	// DELETE /register/:id - ลบข้อมูล
	r.DELETE("/register/:id", func(c *gin.Context) {
		id := c.Param("id")

		// ตรวจสอบว่ามีข้อมูลอยู่หรือไม่
		var register Register
		if err := db.First(&register, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล!"})
			return
		}

		// ลบข้อมูล
		result := db.Delete(&register)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted ข้อมูลสำเร็จ",
		})
	})

	r.Run()

}
