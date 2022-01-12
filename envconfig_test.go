package envconfig

import (
	"testing"
)

type config_test struct {
	I int32 `default:"123"`
}

func TestLoadConfig(t *testing.T) {
	config_interface, _ := LoadConfig(&config_test{})
	realconfig := config_interface.(config_test)
	got := realconfig.I
	if got != 123 {
		t.Errorf("I = %v, 123 expected", got)
	}
}
