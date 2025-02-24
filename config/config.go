package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var apiKey string

func LoadEnv() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("⚠️  Aviso: Arquivo .env não encontrado, usando variáveis do sistema")
    }

    apiKey = os.Getenv("CRYPTO_API_KEY")
}

// GetAPIKey retorna a API Key armazenada
func GetAPIKey() string {
    return apiKey
}