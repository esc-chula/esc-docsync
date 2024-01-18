package config

import (
	"io"
	"os"

	"github.com/esc-chula/esc-docsync/platform/logger"
)

func ReadMap(mapFile string) ([]byte, error) {
	log := logger.GetLogger()

	jsonFile, err := os.Open(mapFile)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	return byteValue, nil
}
