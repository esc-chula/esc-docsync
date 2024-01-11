package config

import (
	"os"
)

type NocoDB struct {
	NocoDBAPIToken string
}

var nocodb = &NocoDB{}

func NocoDBConfig() *NocoDB {
	return nocodb
}

func LoadNocoDBConfig() {
	nocodb.NocoDBAPIToken = os.Getenv("NOCODB_API_TOKEN")
}
