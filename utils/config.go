package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfigs struct {
	MONGOURI, TOKEN_SECRET string
	EncryptionKey        string
	InitVector           string
	PORT                 string
}

var uri, secret, port string
var ok bool

//  encryptionKey, initVector string

func ValidateEnvVars() {
	godotenv.Load()

	uri, ok = os.LookupEnv("MONGOURI")
	if !ok {
		panic("MONGOURI  is required for the server to run")
	}
	secret, ok = os.LookupEnv("TOKEN_SECRET")
	if !ok {
		panic("TOKEN_SECRET is required for the server to run")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8090"
	}

	// encryptionKey, ok = os.LookupEnv("ENCRYPTION_KEY")
	// if !ok {
	// 	panic("ENCRYPTION_KEY is required for the server to run")
	// }

	// initVector, ok = os.LookupEnv("INIT_VECTOR")
	// if !ok {
	// 	panic("INIT_VECTOR is required for the server to run")
	// }
}

var EnvVariable = EnvConfigs{
	MONGOURI:   uri,
	TOKEN_SECRET: secret,
	PORT:       port,
}
