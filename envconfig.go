package envconfig

import (
	"fmt"
	"os"
	"reflect"
)

func LoadConfig(config interface{}) (interface{}, error) {
	ps := reflect.ValueOf(config)
	s := ps.Elem()
	v := reflect.ValueOf(config)
    i := reflect.Indirect(v)
    st := i.Type()
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		toset := s.Field(i)
		if value, exist := os.LookupEnv(field.Name); exist {
			setFieldFromString(field, toset, value)
		} else if def, ok := st.Field(i).Tag.Lookup("default"); ok {
			setFieldFromString(field, toset, def)
		} else {
			return nil, fmt.Errorf("no def or env")
		}
	}
	return s.Interface(), nil
}
