package config

import (
	"context"
	"fmt"
	"inventory-system/models"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

// ConnectDatabase menghubungkan ke PostgreSQL
func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("✅ Berhasil konek ke database")
	DB = database

	DB.AutoMigrate(&models.Item{}, &models.Transaction{})
}

// ConnectRedis menghubungkan ke Redis
func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // atau gunakan strconv.Atoi(os.Getenv("REDIS_DB")) jika ingin dinamis
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Gagal connect ke Redis:", err)
	}

	log.Println("✅ Berhasil konek ke Redis")
}

// Inisialisasi semua koneksi
func Init() {
	ConnectDatabase()
	ConnectRedis()
}
