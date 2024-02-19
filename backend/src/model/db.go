package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB接続とテーブル作成
func ConnectionDB() {
	// DB接続
	dsn := GetDBConfig()
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err)
	}

	// Task型のテーブル作成
	CreateTable(db)
	// *gorm.DB型を *sql.DB型に変換する
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}
	return sqlDB
}

// DBのdsnを取得する
func GetDBConfig() string {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

// Task型のテーブルを作成
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
