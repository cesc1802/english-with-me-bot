package utils

import (
	"log"
	"os"

	"github.com/cesc1802/english-with-me-bot/pkg/statics"
)

func LoadAdaptiveEnvFile() string {
	if len(os.Args) < 2 {
		log.Fatal("Please provide an environment argument: 'dev' or 'prod'")
	}

	serviceEnv := os.Args[1]
	var envFile string
	switch serviceEnv {
	case statics.SERVICE_ENV_DEV:
		envFile = ".env.dev"
	case statics.SERVICE_ENV_PROD:
		envFile = ".env.prod"
	default:
		log.Fatal("Invalid environment. Use 'dev' or 'prod'")
	}

	return envFile
}
