package so

import "os"

type EnvironmentInterface interface {
	GetEnvironmentConfigs() *EnvironmentConfig
}

type Environment struct {
}

func NewEnvironment() EnvironmentInterface {
	return &Environment{}
}

const goEnvironmentKey = "GO_ENVIRONMENT"

func (a *Environment) GetEnvironmentConfigs() *EnvironmentConfig {
	env := &EnvironmentConfig{
		Scope: os.Getenv(goEnvironmentKey),
	}

	return env
}

type EnvironmentConfig struct {
	Scope string
}
