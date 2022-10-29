package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DownloadUrl       string
	FileId            string
	DownloadUrlEnding string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			DownloadUrl:       os.Getenv(DownloadUrl),
			FileId:            os.Getenv(FileId),
			DownloadUrlEnding: os.Getenv(DownloadUrlEnding),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("./../.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
