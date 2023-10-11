package so

import (
	"fmt"
	"os"
)

type Environment interface {
	GetEnv(name string) string
	GetEnvOrDefault(name string, defaultValue string) string
	GetIntEnvOrDefault(name string, defaultValue int) int
	GetEnvironmentConfigs() *EnvironmentConfig
}

type environment struct{}

func NewEnvironment() Environment {
	return &environment{}
}

const goEnvironmentKey = "GO_ENVIRONMENT"

func (a *environment) GetEnvironmentConfigs() *EnvironmentConfig {
	env := &EnvironmentConfig{
		Scope: a.GetEnv(goEnvironmentKey),
	}

	return env
}

func (a *environment) GetEnvOrDefault(name string, defaultValue string) string {
	value := a.GetEnv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func (a *environment) GetIntEnvOrDefault(name string, defaultValue int) int {
	value := a.GetEnv(name)
	if value == "" {
		return defaultValue
	}
	var integerValue int

	fmt.Sscan(value, &integerValue)

	return integerValue
}

func (a *environment) GetEnv(name string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}
	return value
}

type EnvironmentConfig struct {
	Scope string
}
