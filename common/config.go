package common

type ConfigService interface {
	GetConfig() Config
}

type Config struct {
	MaxLength int
}

func GetConfig() Config {
	return Config{MaxLength: 100}
}
