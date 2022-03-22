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

const goEnvironmentValue = "production"

func Test_GetEnvironmentConfigs_Success(t *testing.T) {
	os.Setenv(goEnvironmentKey, goEnvironmentValue)

	configs := NewEnvironment().GetEnvironmentConfigs()

	os.Setenv(goEnvironmentKey, "")

	assert.NotNil(t, configs)
	assert.Equal(t, goEnvironmentValue, configs.Scope)
}
