package loadenvs

import (
	"os"
	"reflect"
)

const environmentTagName = "env"

func LoadEnvironments(applicationConfig interface{}) {
	reflectType := reflect.TypeOf(applicationConfig).Elem()
	reflectValue := reflect.ValueOf(applicationConfig).Elem()

	for i := 0; i < reflectType.NumField(); i++ {
		if reflectValue.Field(i).Kind() == reflect.Struct {
			valueAddr := reflectValue.Field(i).Addr()
			LoadEnvironments(valueAddr.Interface())
		}

		secretTagValue := reflectType.Field(i).Tag.Get(environmentTagName)
		if len(secretTagValue) > 0 {
			value := os.Getenv(secretTagValue)
			reflectValue.Field(i).SetString(value)
		}
	}
}
