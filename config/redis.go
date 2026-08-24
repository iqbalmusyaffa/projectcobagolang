package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// InitRedis membuat dan menguji koneksi ke Redis Server via go-redis/v9.
func InitRedis() *redis.Client {
	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // Gunakan default DB 0
	})

	// Pengujian koneksi dengan PING (timeout 3 detik)
	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := rdb.Ping(pingCtx).Result()
	if err != nil {
		fmt.Printf("⚠️ Warning: Gagal terhubung ke Redis Server (%s:%s): %v. Aplikasi akan berjalan tanpa Redis Caching/Blacklist.\n", host, port, err)
		return nil
	}

	fmt.Println("⚡ Berhasil terhubung ke Redis Server!")
	return rdb
}
