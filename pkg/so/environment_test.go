package so

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

const (
	goEnvironmentValue         = "production"
	testEnvKey                 = "test_env_key"
	testEnvValue               = "test_env_value"
	testDefaultEnvKey          = "test_default_env_key"
	testDefaultEnvValue        = "test_default_env_value"
	testBooleanDefaultEnvKey   = "test_bool_default_env_key"
	testBooleanDefaultEnvValue = "false"
)

func Test_GetEnvironmentConfigs_Success(t *testing.T) {
	os.Setenv(goEnvironmentKey, goEnvironmentValue)

	configs := NewEnvironment().GetEnvironmentConfigs()

	os.Setenv(goEnvironmentKey, "")

	assert.NotNil(t, configs)
	assert.Equal(t, goEnvironmentValue, configs.Scope)
}

func Test_GetEnv_Success(t *testing.T) {
	os.Setenv(testEnvKey, testEnvValue)

	value := NewEnvironment().GetEnv(testEnvKey)

	os.Setenv(testEnvKey, "")

	assert.NotNil(t, value)
	assert.Equal(t, testEnvValue, value)
}

func Test_GetEnvOrDefault_Success(t *testing.T) {
	os.Setenv(testDefaultEnvKey, testDefaultEnvValue)

	value := NewEnvironment().GetEnvOrDefault(testDefaultEnvKey, "")

	os.Setenv(testDefaultEnvKey, "")

	assert.NotNil(t, value)
	assert.Equal(t, testDefaultEnvValue, value)
}

func Test_GetBooleanEnvOrDefault_Success(t *testing.T) {
	os.Setenv(testBooleanDefaultEnvKey, testBooleanDefaultEnvValue)

	value := NewEnvironment().GetBooleanEnvOrDefault(testBooleanDefaultEnvKey, true)

	os.Setenv(testBooleanDefaultEnvKey, "")

	assert.NotNil(t, value)
	assert.False(t, value)
}
