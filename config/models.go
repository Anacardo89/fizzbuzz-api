package config

import "time"

type Config struct {
	Server Server `yaml:"server"`
	Token  Token  `yaml:"token"`
	DB     DB     `yaml:"db"`
	Log    Log    `yaml:"logging"`
	Pag    Pag    `yaml:"pagination"`
	DD     DD     `yaml:"datadog"`
}

type Server struct {
	Port            string        `env:"PORT" envDefault:"8080"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`     // seconds
	WriteTimeout    time.Duration `yaml:"write_timeout"`    // seconds
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"` // seconds
}

type Token struct {
	Secret   string        `env:"TOKEN_SECRET" envDefault:"token-secret"`
	Duration time.Duration `yaml:"duration"` // minutes
}

type DB struct {
	DSN             string        `env:"DB_DSN" envDefault:"postgres://user:pass@localhost:5432/dbname?sslmode=disable"`
	MaxConns        int32         `yaml:"max_conns"`
	MinConns        int32         `yaml:"min_conns"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`  // minutes
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time"` // minutes
}

type Log struct {
	Path       string `env:"LOG_PATH" envDefault:"/fizzbuzz-api/logs"`
	File       string `env:"LOG_FILE" envDefault:"fizzbuzz-api.log"`
	Level      string `env:"LOG_LEVEL" envDefault:"info"`
	MaxSize    int    `yaml:"max_size"` // MB
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"` // days
	Compress   bool   `yaml:"compress"`
}

type Pag struct {
	DefaultPageSize int `yaml:"default_page_size"`
	MaxPageSize     int `yaml:"max_page_size"`
}

type DD struct {
	Env              string  `env:"DD_ENV" envDefault:"dev"`
	Service          string  `env:"DD_SERVICE" envDefault:"fizzbuzz-api"`
	Agent            string  `yaml:"agent"`
	Version          string  `env:"DD_VERSION" envDefault:"0.1.0"`
	TracerPort       string  `env:"DD_TRACER_PORT" envDefault:"8126"`
	TracerSampleRate float64 `yaml:"tracer_sample_rate"` // Percentage
	MetricsPort      string  `env:"DD_METRICS_PORT" envDefault:"8125"`
	MetricsNamespace string  `yaml:"metrics_namespace"`
}
