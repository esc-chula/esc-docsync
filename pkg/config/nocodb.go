package config

import (
	"os"
)

type NocoDB struct {
	NocoDBURL      string
	NocoDBAPIToken string
}

var nocodb = &NocoDB{}

func NocoDBConfig() *NocoDB {
	return nocodb
}

func LoadNocoDBConfig() {
	nocodb.NocoDBURL = os.Getenv("NOCODB_URL")
	nocodb.NocoDBAPIToken = os.Getenv("NOCODB_API_TOKEN")
}
