package configs

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct{
	Server struct{
		Port int `yaml:"port"  env-default:"8080"`
		Host string `yaml:"host" env-default:"0.0.0.0"`
	} `yaml:"server"`
	DB struct{
		Port int `yaml:"port"  env-default:"5432"`
		Host string `yaml:"host" env-default:"localhost"`
		SSLMode string `yaml:"sslmode" env-default:"disable"`
	} `yaml:"db"`
}


func GetConfig() Config {
	var cfg Config
	err := cleanenv.ReadConfig("configs/config.yml",&cfg)
	if err != nil {
		log.Fatalf("Ошибка при чтении конфига: %v",err)
	}
	return cfg
}

func GetDBUrl() string {
	_ = godotenv.Load(".env")
	return os.Getenv("DB_URL")
}
