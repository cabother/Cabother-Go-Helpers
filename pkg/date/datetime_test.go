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

func Test_DatetimeNowFormatDatetime_Success(t *testing.T) {
	data := time.Now()
	responseData := NewDatetime().FormatDatetime(data, "dd-MM-yyyy'T'HH:mm:ssZ")
	assert.NotNil(t, responseData)
}

func Test_FormatDateAndTime_Success(t *testing.T) {
	responseData, formatError := NewDatetime().FormatDate("2021-12-27 19:35:44", "dd-MM-yyyy'T'HH:mm:ssZ")
	assert.Nil(t, formatError)
	assert.NotNil(t, responseData)
}

func Test_FormatDate_Success(t *testing.T) {
	responseData, formatError := NewDatetime().FormatDate("2021-12-27", "dd-MM-yyyy'T'HH:mm:ssZ")
	assert.Nil(t, formatError)
	assert.NotNil(t, responseData)
}

func Test_FormatDateWhitoutT_Success(t *testing.T) {
	responseData, formatError := NewDatetime().FormatDate("2021-12-27", "dd-MM-yyyy HH:mm:ssZ")
	assert.Nil(t, formatError)
	assert.NotNil(t, responseData)
}

func Test_FormatDateWhitoutMilisseconds_Success(t *testing.T) {
	responseData, formatError := NewDatetime().FormatDate("2021-12-27", "dd-MM-yyyy HH:mm:ss")
	assert.Nil(t, formatError)
	assert.NotNil(t, responseData)
}
