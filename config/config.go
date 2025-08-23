package config

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v9"
	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configBytes []byte

func LoadConfig() (*Config, error) {
	cfg := defaultConfig()
	if err := yaml.Unmarshal(configBytes, cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	cfg.Server.ReadTimeout *= time.Second
	cfg.Server.WriteTimeout *= time.Second
	cfg.Server.ShutdownTimeout *= time.Second
	cfg.Token.Duration *= time.Minute
	cfg.DB.MaxConnLifetime *= time.Minute
	cfg.DB.MaxConnIdleTime *= time.Minute
	envs := EnvVars{}
	if err := env.Parse(&envs); err != nil {
		return nil, fmt.Errorf("parsing env: %w", err)
	}
	cfg.Server.Port = envs.Port
	cfg.DB.DSN = envs.DBDSN
	cfg.Log.Path = envs.LogPath
	cfg.Log.File = envs.LogFile
	cfg.Log.Level = envs.LogLevel
	log.Printf("final cfg: %+v", cfg)
	return cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			ReadTimeout:     5,  // seconds
			WriteTimeout:    10, // seconds
			ShutdownTimeout: 15, // seconds
		},
		Token: TokenConfig{
			Duration: 60 * time.Minute, // minutes
		},
		DB: DBConfig{
			MaxConns:        10,
			MinConns:        2,
			MaxConnLifetime: 30, // minutes
			MaxConnIdleTime: 5,  // minutes
		},
		Log: LogConfig{
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
		},
		Pag: PagConfig{
			DefaultPageSize: 20,
			MaxPageSize:     200,
		},
	}
}
