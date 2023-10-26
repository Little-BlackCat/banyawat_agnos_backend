package main

import (
	"database/sql"
	"log"
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type UserData struct {
	Password string `json:"init_password"`
}

var db *sql.DB

func main() {
	// Connect database PostgreSQL
	connStr := "user=postgres dbname=postgres password=password sslmode=disable host=db port=5432"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal(err)
	}

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)
	
	// Initialize the Gin router
	router := gin.Default()

	// GET testing
	router.GET("/testing", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	
	// POST เส้นทางสำหรับรับข้อมูลจากผู้ใช้
	router.POST("/api/check_password", func(ctx *gin.Context) {
		var userData UserData

		// ใช้ Bind เพื่อแปลงข้อมูล JSON ที่ผู้ใช้ส่งมาเป็น struct
		if err := ctx.BindJSON(&userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		action, isStrong := CheckStrongPassword(userData.Password)

		// Store request & response into database
		saveRequestResponse(userData.Password, action, isStrong)

		ctx.JSON(http.StatusOK, gin.H {
			"message": "save user's password",
			"num_of_steps": action,
			"isStrong": isStrong,
		})
	})

	// Start the server
	router.Run(":8080")

	// Close the database connection after saving
	db.Close()
}

func CheckStrongPassword(password string) (int, bool) {

	actions := 0
	isStrong := true

	// Criteria 1: Password length >=6, < 20 characters.
	if len(password) < 6 || len(password) >= 20 {
		actions++
		isStrong = false
	}

	lowercase := false
	uppercase := false
	digit := false
	repeatingCount := 1

	for i := 0; i < len(password); i++ {
		// Criteria 3: Does not contain 3 repeating characters in a row e.g. 11123.
		if i > 0 && password[i] == password[i-1] {
			repeatingCount++
		} else {
			repeatingCount = 1
		}

		if repeatingCount >= 3 {
			actions++
			isStrong = false
			break
		}

		if isLetter(password[i]) {
			// Criteria 2: Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit.
			if isLowercase(password[i]) {
				lowercase = true
			} else if isUppercase(password[i]) {
				uppercase = true
			}
		} else if isDigit(password[i]) {
			digit = true
		}
	}

	// Criteria 2: Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit.
	if !lowercase || !uppercase || !digit {
		actions++
		isStrong = false
	}

	return actions, isStrong
}


func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isLowercase(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func isUppercase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func createSchema() {
	_, err := db.Exec(`CREATE SCHEMA IF NOT EXISTS public;`)
	if err != nil {
		log.Fatal("Error creating schema:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS public.password_log (
			id serial PRIMARY KEY,
			password text,
			num_of_steps int,
			is_strong boolean
		);
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func saveRequestResponse(password string, action int, isStrong bool) {
	createSchema()

	_, err := db.Exec("INSERT INTO password_log (password, num_of_steps, is_strong) VALUES ($1, $2, $3)", password, action, isStrong)
	if err != nil {
		log.Println("Error saving request and response:", err)
	}

	
}

