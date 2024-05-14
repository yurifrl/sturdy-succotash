package config

// Config interface provides configuration for the Consumer
type Config interface {
	Food() string
}

type EnvConfig struct{}

func (e *EnvConfig) Food() string {
	return ""
}
