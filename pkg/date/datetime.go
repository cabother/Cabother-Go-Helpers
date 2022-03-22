package date

import (
	"strings"
	"time"
)

type DatetimeInterface interface {
	// FormatDate is a function to format a received date according to the custom parameter 'format'
	FormatDate(date time.Time, format string) string

	// DatetimeNow is a function to get current date and time and format according to custom parameter 'format'
	GetDateNow(format string) string
}

type Datetime struct{}

func NewDatetime() DatetimeInterface {
	return &Datetime{}
}

func (a *Datetime) FormatDate(date time.Time, format string) string {
	return date.Format(a.replaceFormat(format))
}

func (a *Datetime) GetDateNow(format string) string {
	return time.Now().Format(a.replaceFormat(format))
}

func (a *Datetime) replaceFormat(format string) string {
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "yy", "06")
	format = strings.ReplaceAll(format, "MMMM", "January")
	format = strings.ReplaceAll(format, "MMM", "Jan")
	format = strings.ReplaceAll(format, "MM", "01")
	format = strings.ReplaceAll(format, "dd", "02")
	format = strings.ReplaceAll(format, "d ", "2 ")
	format = strings.ReplaceAll(format, "HH", "15")
	format = strings.ReplaceAll(format, "mm", "04")
	format = strings.ReplaceAll(format, "ssZ", "05-0700")
	format = strings.ReplaceAll(format, "ss", "05")
	format = strings.ReplaceAll(format, "EEEE", "Monday")
	format = strings.ReplaceAll(format, "EEE", "Mon")

	return format
}
