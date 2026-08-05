package config

import "os"

func AppURL() string {
	return os.Getenv("APP_URL")
}
