package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host           string
	Port           int
	Debug          bool
	ReadTimeout    time.Duration
	RequestTimeout time.Duration
	AllowedOrigins string
	SecretKey      string
}

var app = &App{}

func AppConfig() *App {
	return app
}

func LoadAppConfig() {
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second
	timeOut, _ = strconv.Atoi(os.Getenv("APP_REQUEST_TIMEOUT"))
	app.RequestTimeout = time.Duration(timeOut) * time.Second
	app.AllowedOrigins = os.Getenv("APP_ALLOWED_ORIGINS")
	app.SecretKey = os.Getenv("APP_SECRET_KEY")
}
