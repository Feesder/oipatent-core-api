package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"database"`
	Token      `yaml:"token"`
	Security   `yaml:"security"`
	Cookie     `yaml:"cookie"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Database struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

type Token struct {
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
	AccessTTL     time.Duration `yaml:"access_ttl"`
	RefreshTTL    time.Duration `yaml:"refresh_ttl"`
	Issuer        string        `yaml:"issuer"`
}

type Security struct {
	Salt string `yaml:"salt"`
}

type Cookie struct {
	Path     string `yaml:"path"`
	Domain   string `yaml:"domain"`
	Secure   bool   `yaml:"secure"`
	HttpOnly bool   `yaml:"http_only"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
