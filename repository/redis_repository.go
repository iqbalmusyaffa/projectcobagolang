package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisRepository mendefinisikan interface operasi Caching, Token Blacklisting, & Refresh Token di Redis.
type RedisRepository interface {
	BlacklistToken(token string, expiration time.Duration) error
	IsTokenBlacklisted(token string) bool
	StoreRefreshToken(userID uint, refreshToken string, expiration time.Duration) error
	GetRefreshToken(userID uint) (string, error)
	DeleteRefreshToken(userID uint) error
	SetCache(key string, value string, expiration time.Duration) error
	GetCache(key string) (string, error)
	DeleteCache(key string) error
}

type redisRepository struct {
	rdb *redis.Client
	ctx context.Context
}

// NewRedisRepository membuat instance RedisRepository.
func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return &redisRepository{
		rdb: rdb,
		ctx: context.Background(),
	}
}

// BlacklistToken menyimpan token JWT yang di-logout ke Redis Blacklist dengan Expiration (TTL).
func (r *redisRepository) BlacklistToken(token string, expiration time.Duration) error {
	if r.rdb == nil {
		return nil
	}
	key := fmt.Sprintf("blacklist:%s", token)
	return r.rdb.Set(r.ctx, key, "logout", expiration).Err()
}

// IsTokenBlacklisted mengecek apakah token JWT terdaftar di Redis Blacklist.
func (r *redisRepository) IsTokenBlacklisted(token string) bool {
	if r.rdb == nil {
		return false
	}
	key := fmt.Sprintf("blacklist:%s", token)
	val, err := r.rdb.Get(r.ctx, key).Result()
	return err == nil && val == "logout"
}

// StoreRefreshToken menyimpan Refresh Token di Redis dengan key refresh_token:<userID>.
func (r *redisRepository) StoreRefreshToken(userID uint, refreshToken string, expiration time.Duration) error {
	if r.rdb == nil {
		return nil
	}
	key := fmt.Sprintf("refresh_token:%d", userID)
	return r.rdb.Set(r.ctx, key, refreshToken, expiration).Err()
}

// GetRefreshToken mengambil Refresh Token dari Redis berdasarkan userID.
func (r *redisRepository) GetRefreshToken(userID uint) (string, error) {
	if r.rdb == nil {
		return "", fmt.Errorf("redis tidak aktif")
	}
	key := fmt.Sprintf("refresh_token:%d", userID)
	return r.rdb.Get(r.ctx, key).Result()
}

// DeleteRefreshToken menghapus Refresh Token dari Redis berdasarkan userID.
func (r *redisRepository) DeleteRefreshToken(userID uint) error {
	if r.rdb == nil {
		return nil
	}
	key := fmt.Sprintf("refresh_token:%d", userID)
	return r.rdb.Del(r.ctx, key).Err()
}

// SetCache menyimpan nilai string ke Redis Cache berdasarkan Key.
func (r *redisRepository) SetCache(key string, value string, expiration time.Duration) error {
	if r.rdb == nil {
		return nil
	}
	return r.rdb.Set(r.ctx, key, value, expiration).Err()
}

// GetCache mengambil nilai dari Redis Cache berdasarkan Key.
func (r *redisRepository) GetCache(key string) (string, error) {
	if r.rdb == nil {
		return "", fmt.Errorf("redis tidak aktif")
	}
	return r.rdb.Get(r.ctx, key).Result()
}

// DeleteCache menghapus data dari Redis Cache berdasarkan Key.
func (r *redisRepository) DeleteCache(key string) error {
	if r.rdb == nil {
		return nil
	}
	return r.rdb.Del(r.ctx, key).Err()
}
