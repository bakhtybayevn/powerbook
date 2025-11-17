package config

type AppConfig struct {
    Environment string `mapstructure:"environment"`
    Port        int    `mapstructure:"port"`
}

type PostgresConfig struct {
    DSN string `mapstructure:"dsn"`
}

type RedisConfig struct {
    Addr string `mapstructure:"addr"`
}

type TelegramConfig struct {
    Token string `mapstructure:"token"`
}

type Config struct {
    App      AppConfig      `mapstructure:"app"`
    Postgres PostgresConfig `mapstructure:"postgres"`
    Redis    RedisConfig    `mapstructure:"redis"`
    Telegram TelegramConfig `mapstructure:"telegram"`
}
