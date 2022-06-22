package date

import (
	"fmt"

	pkgDate "github.com/cabother/cabother-go-helpers/pkg/date"
)

type DateTimeExampleInterface interface {
	GetCurrentDateTime() error
}

type DateTimeExample struct{}

func New() DateTimeExampleInterface {
	return &DateTimeExample{}
}

func (a *DateTimeExample) GetCurrentDateTime() error {
	instance := pkgDate.NewDatetime()
	currentDateTime := instance.GetDateNow("dd/MM/yyyyT00:00:00")
	//... your code here

	fmt.Println(currentDateTime)
	return nil
}
