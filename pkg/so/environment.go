package so

import "os"

type EnvironmentInterface interface {
	GetEnv(name string) string
	GetEnvOrDefault(name string, defaultValue string) string
	GetEnvironmentConfigs() *EnvironmentConfig
}

type Environment struct{}

func NewEnvironment() EnvironmentInterface {
	return &Environment{}
}

const goEnvironmentKey = "GO_ENVIRONMENT"

func (a *Environment) GetEnvironmentConfigs() *EnvironmentConfig {
	env := &EnvironmentConfig{
		Scope: a.GetEnv(goEnvironmentKey),
	}

	return env
}

func (a *Environment) GetEnvOrDefault(name string, defaultValue string) string {
	value := a.GetEnv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func (a *Environment) GetEnv(name string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}
	return value
}

type EnvironmentConfig struct {
	Scope string
}
