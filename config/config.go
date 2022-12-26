package config

type Config struct {
	BaseURL string `envconfig:"BASE_URL"`
	ApiKey  string `envconfig:"X_API_KEY"`
}
