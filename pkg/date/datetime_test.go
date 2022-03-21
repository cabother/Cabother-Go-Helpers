package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DatetimeNowFormat1_Success(t *testing.T) {
	data := NewDatetime().DatetimeNow(Format1)
	assert.NotNil(t, data)
}

func Test_DatetimeNowFormat2_Success(t *testing.T) {
	data := NewDatetime().DatetimeNow(Format2)
	assert.NotNil(t, data)
}
