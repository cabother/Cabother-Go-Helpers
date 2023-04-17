package health

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Health_Instance(t *testing.T) {
	// Given
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	h := handler{}

	// When
	h.Health(c)

	// Then
	assert.NotNil(t, h)
}
