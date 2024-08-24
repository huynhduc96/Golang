package main

import (
	"database/sql"
	"fmt"
	"log"
	"social/internal/connector"
	"social/internal/constant"
	"social/internal/utils"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	// Seed MySQL Database
	// seedMySqlDatabase()

	// Seed Redis Database
	seedRedisDatabase()
}

func seedRedisDatabase() {
	log.Printf("Start Seed Cache \n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cacheClient, err := connector.CreateRedisClient()
	if err != nil {
		log.Fatal(err)
		log.Printf("err%+v\n", err)
	}
	defer cacheClient.Close()

	for i := 1; i < 5; i++ {
		sessionID := fmt.Sprintf("session-username%d", i)
		cacheClient.Set(sessionID, strconv.Itoa(i), constant.TtlSeed).Err()
	}
	log.Printf("Done Seed Cache \n")

}

func seedMySqlDatabase() {
	log.Printf("Start Seed Data \n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Database connection
	db, err := connector.CreateDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	// seed Users
	seedUsers(db)

	// seed Follower
	seedFollows(db)

	// seed Posts
	seedPosts(db)

	log.Printf("Done Seed Data \n")
}

func seedUsers(db *sql.DB) {

	log.Printf("Start seed users\n")

	stmtIn, err := db.Prepare("INSERT INTO user (first_name, last_name, user_name, email, salt, hashed_password, dob) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIn.Close()

	salt := utils.GenRandomSalt(constant.SaltSize)

	hashedPass := utils.HashPassword("123456", salt)

	for i := 0; i <= 1000000; i++ {
		firstName := fmt.Sprintf("first name %d", i)
		lastName := fmt.Sprintf("last name %d", i)
		userName := fmt.Sprintf("username%d", i)
		email := fmt.Sprintf("example%d@example.com", i)

		_, errQuery := stmtIn.Exec(firstName, lastName, userName, email, salt, hashedPass, "2014-04-24")

		if errQuery != nil {
			log.Printf("Error%+v\n", errQuery)
		}
	}

	log.Printf("Done seed users\n")
}

func seedFollows(db *sql.DB) {

	log.Printf("Start seed Follows\n")

	stmtIn, err := db.Prepare("INSERT INTO `user_user` (fk_user_id, fk_follower_id) VALUES (?, ?)")
	defer stmtIn.Close()

	if err != nil {
		log.Fatal(err)
	}

	// create user id 0 --> 100 have 100000 followers
	for i := 0; i <= 100; i++ {
		for j := i + 1; j <= 100000+1; j++ {
			_, err = stmtIn.Exec(i, j)
		}
	}

	// create user id 101 -->  500 have 200 followers
	for i := 101; i <= 500; i++ {
		for j := i + 1; j <= 200+1; j++ {
			_, err = stmtIn.Exec(i, j)
		}
	}

	log.Printf("Done seed Follows\n")
}

func seedPosts(db *sql.DB) {

	log.Printf("Start seed Posts\n")

	stmtIn, err := db.Prepare("INSERT INTO post (fk_user_id, content_text) VALUES (?, ?)")
	defer stmtIn.Close()

	if err != nil {
		log.Fatal(err)
	}

	// 100 first user have 3 posts
	for i := 0; i <= 100; i++ {

		for j := 1; j <= 3; j++ {

			contentTextPost := fmt.Sprintf("Content text from user %d : post number %d", i, j)

			_, errQuery := stmtIn.Exec(
				i,
				contentTextPost,
			)

			if errQuery != nil {
				log.Printf("Error%+v\n", errQuery)
			}
		}

	}

	// 10000 next user have 1 post

	for i := 101; i <= 10000; i++ {
		contentTextPost := fmt.Sprintf("Content text from user %d", i)

		_, errQuery := stmtIn.Exec(
			i,
			contentTextPost,
		)

		if errQuery != nil {
			log.Printf("Error%+v\n", errQuery)
		}
	}

	log.Printf("Done seed Posts\n")
}
