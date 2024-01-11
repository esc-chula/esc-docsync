package config

type App struct {
}

var app = &App{}

func AppConfig() *App {
	return app
}

func LoadAppConfig() {
}
