package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var redisClient redis.Conn

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	ConnectDatabases()
}

func ConnectDatabases() {
	db = connectPostgres()
	// Skipping redis connection for now
	// redisClient = connectRedis()
}

func connectPostgres() *gorm.DB {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	sslMode := os.Getenv("sslmode")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName, sslMode, password)

	var conn *gorm.DB
	var err error
	retryLimit := time.Now().Add(1 * time.Minute)

	for {
		conn, err = gorm.Open("postgres", dbUri)
		if err == nil {
			break
		}
		fmt.Println("Failed to connect to database. Retrying...")
		time.Sleep(5 * time.Second)
		if time.Now().After(retryLimit) {
			fmt.Println("Exceeded retry limit. Continuing without database connection...")
			return nil
		}
	}
	fmt.Println("Connected to PostgreSQL:", dbUri)
	return conn
}

func connectRedis() redis.Conn {
	redisConnection, err := redis.Dial("tcp", os.Getenv("redis_host")+":6379")
	if err != nil {
		fmt.Println("Failed to connect to Redis. Continuing without Redis connection...")
		return nil
	}
	fmt.Println("Connected to Redis")
	return redisConnection
}

func GetDB() *gorm.DB {
	return db
}

func GetRedis() redis.Conn {
	return redisClient
}
