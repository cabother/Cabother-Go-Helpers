package date

import (
	"time"
)

type DatetimeInterface interface {
	// DatetimeNow is a function to get datetime now and format by param 'format'
	DatetimeNow(datetimeType DatetimeType) string
}

type Datetime struct {
}

func NewDatetime() DatetimeInterface {
	return &Datetime{}
}

const (
	ddMMyyyyHHmmss  = "02-01-2006 15:04:05"
	ddMMyyyyHHmmss2 = "02/01/2006 15:04:05"

	yyyyMMddTHHmmss        = "2006-01-02 15:04:05"
	yyyyMMddTHHmmssZ       = "2006-01-02T15:04:05-0700"
	twoJan2006150405       = "2 Jan 2006 15:04:05"
	twoJan20061504         = "2 Jan 2006 15:04"
	monTwoJan2006150405MST = "Mon, 2 Jan 2006 15:04:05 MST"
)

type DatetimeType string

const (
	// format {dd-MM-yyyy HH:mm:ss}
	//
	// example: {02-01-2006 15:04:05}
	Format1 DatetimeType = "dd-MM-yyyy HH:mm:ss"
	// format {dd/MM/yyyy HH:mm:ss}
	//
	// example: {02/01/2006 15:04:05}
	Format2 DatetimeType = "dd/MM/yyyy HH:mm:ss"
	// format {yyyy-MM-ddTHH:mm:ss}
	//
	// example: {2006-01-02T15:04:05}
	Format10 DatetimeType = "yyyy-MM-ddTHH:mm:ss"
	// format {yyyy-MM-ddTHH:mm:ssZ}
	//
	// example: {2006-01-02T15:04:05-0700}
	Format11 DatetimeType = "yyyy-MM-ddTHH:mm:ssZ"
	// format {d MMM yyyy HH:mm:ss}
	//
	// example: {2 Jan 2006 15:04:05}
	Format12 DatetimeType = "d MMM yyyy HH:mm:ss"
	// format {d MMM yyyy HH:mm}
	//
	// example: {2 Jan 2006 15:04}
	Format13 DatetimeType = "d MMM yyyy HH:mm"
	// format {EEE, d MMM yyyy HH:mm:ss z}
	//
	// example: {Mon, 2 Jan 2006 15:04:05 MST}
	Format14 DatetimeType = "EEE, d MMM yyyy HH:mm:ss z"
)

func (a *Datetime) DatetimeNow(datetimeType DatetimeType) string {
	switch datetimeType {
	case Format1:
		return time.Now().Format(ddMMyyyyHHmmss)
	case Format2:
		return time.Now().Format(ddMMyyyyHHmmss2)
	case Format10:
		return time.Now().Format(yyyyMMddTHHmmssZ)
	case Format11:
		return time.Now().Format(yyyyMMddTHHmmssZ)
	case Format12:
		return time.Now().Format(twoJan2006150405)
	case Format13:
		return time.Now().Format(twoJan20061504)
	case Format14:
		return time.Now().Format(monTwoJan2006150405MST)
	}
	return time.Now().Format("2006-01-02 15:04:05")
}
