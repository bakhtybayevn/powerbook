package config

import (
    "github.com/spf13/viper"
    "fmt"
)

func Load() (*Config, error) {
    v := viper.New()
    v.SetConfigName("config")
    v.SetConfigType("yaml")
    v.AddConfigPath(".")
    v.AddConfigPath("./config")
    v.AutomaticEnv()

    // defaults
    v.SetDefault("app.environment", "development")
    v.SetDefault("app.port", 8080)
    v.SetDefault("postgres.dsn", "postgres://user:pass@localhost:5432/powerbook?sslmode=disable")
    v.SetDefault("redis.addr", "localhost:6379")
    v.SetDefault("telegram.token", "")

    if err := v.ReadInConfig(); err != nil {
        // not fatal - will use defaults and env vars
        fmt.Println("config file not found, using defaults and environment variables")
    }

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
