package envconfig

import (
	"fmt"
	"reflect"
	"time"
)

func setFieldFromString(field reflect.StructField, toset reflect.Value, value string) error {
	switch field.Type.Kind() {
	case reflect.String:
		toset.SetString(value)
		return nil
	case reflect.Bool:
		bool, err := parseBool(value)
		toset.SetBool(bool)
		return err
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := parseInt64(value)
		toset.SetInt(i)
		return err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := parseUint64(value)
		toset.SetUint(i)
		return err
	case reflect.Float32, reflect.Float64:
		f, err := parseFloat64(value)
		toset.SetFloat(f)
		return err
	default:
		switch field.Type {
		case reflect.TypeOf(time.Duration(0)):
			d, err := parseDuration(value)
			toset.Set(reflect.ValueOf(d))
			return err
		}
	}
	return fmt.Errorf("invalid type: %v", field.Type)
}
