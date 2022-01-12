package envconfig

import (
	"os"
	"testing"
)

type config_test struct {
	I32        int32  `default:"123"`
	SomeString string `name:"CUSTOM_NAME"`
	F32Value   float32
}

func TestLoadConfig(t *testing.T) {
	os.Setenv("CUSTOM_NAME", "SomeString_VALUE")
	os.Setenv("F32_VALUE", "123.456")

	config_interface, err := LoadConfig(&config_test{})
	if err != nil {
		t.Error(err)
	}

	c := config_interface.(config_test)

	if c.I32 != 123 {
		t.Errorf("I32 = %v, 123 expected", c.I32)
	}

	if c.SomeString != "SomeString_VALUE" {
		t.Errorf("SomeString = %v, SomeString_VALUE expected", c.SomeString)
	}

	if c.F32Value != 123.456 {
		t.Errorf("F32 = %v, 123.456 expected", c.F32Value)
	}
}
