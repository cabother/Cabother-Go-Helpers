package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DatetimeNowFormat1_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("dd/MM/yyyy")
	assert.NotNil(t, data)
}

func Test_DatetimeNowFormat2_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("dd/MM/yyyyTHH:mm:ss")
	assert.NotNil(t, data)
}

func Test_DatetimeNowFormat3_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("yyyy-MM-dd'T'HH:mm:ssZ")
	assert.NotNil(t, data)
}

func Test_DatetimeNowFormat4_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("dd-MM-yyyy'T'HH:mm:ssZ")
	assert.NotNil(t, data)
}

func Test_DateFormat1_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("dd/MM/yyyy")
	assert.NotNil(t, data)
}

func Test_DateFormat2_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("MMMM dd, yyyy")
	assert.NotNil(t, data)
}

func Test_HourFormat1_Success(t *testing.T) {
	data := NewDatetime().GetDateNow("HH:mm:ss")
	assert.NotNil(t, data)
}

func Test_DatetimeNowFormatDate_Success(t *testing.T) {
	data := time.Now()
	responseData := NewDatetime().FormatDate(data, "dd-MM-yyyy'T'HH:mm:ssZ")
	assert.NotNil(t, responseData)
}
