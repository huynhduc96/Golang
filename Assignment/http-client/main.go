package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}
	dbURL := os.Getenv("DATABASE_URL")
	return dbURL, nil
}

func main() {
	// Load the .env file
	dbURL, err := LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// check connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// // Tạo bảng
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	//     id INT AUTO_INCREMENT PRIMARY KEY,
	//     name VARCHAR(255) NOT NULL,
	//     email VARCHAR(255) UNIQUE NOT NULL
	// )`)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tạo bảng users thành công!")

	// // Thêm dữ liệu
	// _, err = db.Exec(`INSERT INTO users (name, email) VALUES ('John Doe', 'john.doe@example.com')`)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Thêm dữ liệu thành công!")

	// Get data
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Danh sách users:")
	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// handle error
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
