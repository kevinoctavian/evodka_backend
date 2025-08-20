package config

import (
	"os"
	"strconv"
	"time"
)

// App holds the App configuration
type App struct {
	Host        string
	Port        int
	Debug       bool
	ReadTimeout time.Duration

	// JWT Conf
	JWTAccessKey                   string
	JWTAccessKeyExpireMinutesCount int
	JWTRefreshKey                  string
	JWTRefreshKeyExpireHourCount   int
}

var app = &App{}

// AppCfg returns the default App configuration
func AppCfg() *App {
	return app
}

// LoadApp loads App configuration
func LoadApp() {
	app.Host = os.Getenv("HOST")
	if app.Host == "" {
		app.Host = os.Getenv("APP_HOST")
	}

	var err error
	app.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		app.Port, err = strconv.Atoi(os.Getenv("APP_PORT")) // Default port if not set
		if err != nil {
			app.Port = 8000 // Fallback to a default port
		}
	}

	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second

	app.JWTAccessKey = os.Getenv("JWT_ACCESS_KEY")
	app.JWTAccessKeyExpireMinutesCount, _ = strconv.Atoi(os.Getenv("JWT_ACCESS_KEY_EXPIRE_MINUTES_COUNT"))
	app.JWTRefreshKey = os.Getenv("JWT_REFRESH_KEY")
	app.JWTRefreshKeyExpireHourCount, _ = strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOUR_COUNT"))
}
