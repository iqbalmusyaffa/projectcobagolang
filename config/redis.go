package config

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// InitRedis membuat client Redis dan menguji koneksi ke Redis Server.
//
// Jika Redis tidak tersedia, fungsi mengembalikan nil sehingga
// aplikasi tetap dapat berjalan tanpa Redis.
func InitRedis() *redis.Client {
	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "")
	db := getEnvInt("REDIS_DB", 0)

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,

		// Batas waktu koneksi.
		DialTimeout: 5 * time.Second,

		// Batas waktu operasi baca.
		ReadTimeout: 3 * time.Second,

		// Batas waktu operasi tulis.
		WriteTimeout: 3 * time.Second,

		// Connection Pool.
		PoolSize:     10,
		MinIdleConns: 2,
	})

	// Context dengan timeout untuk testing koneksi.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	// Test koneksi Redis.
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		fmt.Printf(
			"⚠️ Warning: Gagal terhubung ke Redis Server (%s): %v\n",
			addr,
			err,
		)

		// Tutup client jika koneksi gagal.
		if closeErr := rdb.Close(); closeErr != nil {
			fmt.Printf(
				"⚠️ Warning: Gagal menutup Redis client: %v\n",
				closeErr,
			)
		}

		fmt.Println(
			"ℹ️ Aplikasi akan berjalan tanpa Redis Caching/Blacklist.",
		)

		return nil
	}

	fmt.Printf(
		"⚡ Berhasil terhubung ke Redis Server: %s\n",
		addr,
	)

	return rdb
}

// getEnvInt mengambil environment variable dalam bentuk integer.
func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
