package so

import (
	"fmt"
	"os"
)

type EnvironmentInterface interface {
	GetEnv(name string) string
	GetEnvOrDefault(name string, defaultValue string) string
	GetIntEnvOrDefault(name string, defaultValue int) int
	GetBooleanEnvOrDefault(name string, defaultValue bool) bool
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

func (a *Environment) GetIntEnvOrDefault(name string, defaultValue int) int {
	value := a.GetEnv(name)
	if value == "" {
		return defaultValue
	}
	var integerValue int

	fmt.Sscan(value, &integerValue)

	return integerValue
}

func (a *Environment) GetEnv(name string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}
	return value
}

func (a *Environment) GetBooleanEnvOrDefault(name string, defaultValue bool) bool {
	value := a.GetEnv(name)
	if value == "" {
		return defaultValue
	}
	var booleanValue bool

	fmt.Sscan(value, &booleanValue)

	return booleanValue
}

type EnvironmentConfig struct {
	Scope string
}
