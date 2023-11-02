package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis"
	"time"
)

var redisClient *redis.Client
var db *sql.DB

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:1234@localhost:5433/event_calendar_test")
	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return
	}
}

func getDataFromPostgreSQL() ([]byte, error) {
	query := "SELECT data FROM events WHERE id = $1"
	row := db.QueryRow(query, 1)

	var data string
	if err := row.Scan(&data); err != nil {
		return nil, err
	}

	return []byte(data), nil
}

func getOrSetData(key string, expiration time.Duration) ([]byte, error) {
	cachedData, err := redisClient.Get(key).Bytes()
	if err == redis.Nil {
		data, err := getDataFromPostgreSQL()
		if err != nil {
			return nil, err
		}

		err = redisClient.Set(key, data, expiration).Err()
		if err != nil {
			return nil, err
		}

		return []byte(data), nil
	} else if err != nil {
		return nil, err
	}

	return []byte(cachedData), nil
}

func main() {
	cacheKey := "event1"
	data, err := getOrSetData(cacheKey, 30*time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}
}
