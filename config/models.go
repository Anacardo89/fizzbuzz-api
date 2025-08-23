package config

import "time"

// ENV

type EnvVars struct {
	Port        string `env:"PORT" envDefault:"8080"`
	TokenSecret string `env:"JWT_SECRET" envDefault:"token-secret"`
	DBDSN       string `env:"DB_DSN" envDefault:"postgres://user:pass@localhost:5432/dbname?sslmode=disable"`
	LogPath     string `env:"LOG_PATH" envDefault:"/fizzbuzz-api/logs"`
	LogFile     string `env:"LOG_FILE" envDefault:"fizzbuzz-api.log"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
}

// YAML

type Config struct {
	Server ServerConfig `yaml:"server"`
	Token  TokenConfig  `yaml:"token"`
	DB     DBConfig     `yaml:"db"`
	Log    LogConfig    `yaml:"logging"`
	Pag    PagConfig    `yaml:"pagination"`
}

type ServerConfig struct {
	Port            string
	ReadTimeout     time.Duration `yaml:"read_timeout"`     // seconds
	WriteTimeout    time.Duration `yaml:"write_timeout"`    // seconds
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"` // seconds
}

type TokenConfig struct {
	Secret   string
	Duration time.Duration `yaml:"duration"` // minutes
}

type DBConfig struct {
	DSN             string
	MaxConns        int32         `yaml:"max_conns"`
	MinConns        int32         `yaml:"min_conns"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`  // minutes
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time"` // minutes
}

type LogConfig struct {
	Path       string
	File       string
	Level      string
	MaxSize    int  `yaml:"max_size"`
	MaxBackups int  `yaml:"max_backups"`
	MaxAge     int  `yaml:"max_age"`
	Compress   bool `yaml:"compress"`
}

type PagConfig struct {
	DefaultPageSize int `yaml:"default_page_size"`
	MaxPageSize     int `yaml:"max_page_size"`
}
