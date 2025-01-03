package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kazuyainoue0124/go-rest-api/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO tasks (id, title, description, created_at, updated_at)
		VALUES
				(?, ?, ?, ?, ?),
				(?, ?, ?, ?, ?),
				(?, ?, ?, ?, ?),
				(?, ?, ?, ?, ?),
				(?, ?, ?, ?, ?)
		`,
		// 1行目
		1,
		"タイトル1",
		"説明1",
		time.Now(),
		time.Now(),

		// 2行目
		2,
		"タイトル2",
		"説明2",
		time.Now(),
		time.Now(),

		// 3行目
		3,
		"タイトル3",
		"説明3",
		time.Now(),
		time.Now(),

		// 4行目
		4,
		"タイトル4",
		"説明4",
		time.Now(),
		time.Now(),

		// 5行目
		5,
		"タイトル5",
		"説明5",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	fmt.Println("Seed data inserted!")
}
