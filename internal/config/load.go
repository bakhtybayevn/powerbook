package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	v := viper.New()

	// CONFIG FILE
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// LOCAL dev paths
	v.AddConfigPath("./internal/config")
	v.AddConfigPath(".")

	// DOCKER paths
	v.AddConfigPath("/app/internal/config")

	// Try read
	if err := v.ReadInConfig(); err != nil {
		log.Println("config.yaml NOT found, ENV only:", err)
	} else {
		log.Println("config.yaml loaded successfully!")
	}

	// ENV override
	v.AutomaticEnv()

	// bind helper
	bind := func(key, env string) {
		if err := v.BindEnv(key, env); err != nil {
			panic(err)
		}
	}

	// APP
	bind("app.environment", "APP_ENV")
	bind("app.port", "APP_PORT")

	// DATABASE
	bind("database.host", "POSTGRES_HOST")
	bind("database.port", "POSTGRES_PORT")
	bind("database.user", "POSTGRES_USER")
	bind("database.password", "POSTGRES_PASSWORD")
	bind("database.name", "POSTGRES_DB")
	bind("database.sslmode", "POSTGRES_SSLMODE")

	// REDIS
	bind("redis.host", "REDIS_HOST")
	bind("redis.port", "REDIS_PORT")
	bind("redis.password", "REDIS_PASSWORD")
	bind("redis.tls", "REDIS_TLS")

	// JWT
	bind("jwt.secret", "JWT_SECRET")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}

	return &cfg, nil
}
