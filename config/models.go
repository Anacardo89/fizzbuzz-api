package config

import "time"

type Config struct {
	Server ServerConfig `yaml:"server"`
	Token  TokenConfig  `yaml:"token"`
	DB     DBConfig     `yaml:"db"`
	Log    LogConfig    `yaml:"logging"`
	Pag    PagConfig    `yaml:"pagination"`
}

type ServerConfig struct {
	Port            string        `env:"PORT" envDefault:"8080"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`     // seconds
	WriteTimeout    time.Duration `yaml:"write_timeout"`    // seconds
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"` // seconds
}

type TokenConfig struct {
	Secret   string        `env:"TOKEN_SECRET" envDefault:"token-secret"`
	Duration time.Duration `yaml:"duration"` // minutes
}

type DBConfig struct {
	DSN             string        `env:"DB_DSN" envDefault:"postgres://user:pass@localhost:5432/dbname?sslmode=disable"`
	MaxConns        int32         `yaml:"max_conns"`
	MinConns        int32         `yaml:"min_conns"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`  // minutes
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time"` // minutes
}

type LogConfig struct {
	Path       string `env:"LOG_PATH" envDefault:"/fizzbuzz-api/logs"`
	File       string `env:"LOG_FILE" envDefault:"fizzbuzz-api.log"`
	Level      string `env:"LOG_LEVEL" envDefault:"info"`
	MaxSize    int    `yaml:"max_size"` // MB
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"` // days
	Compress   bool   `yaml:"compress"`
}

type PagConfig struct {
	DefaultPageSize int `yaml:"default_page_size"`
	MaxPageSize     int `yaml:"max_page_size"`
}
