package envconfig

import "testing"

func TestFormatNameAsEnv(t *testing.T) {
	got := formatNameAsEnv("PascalCase")
	if got != "PASCAL_CASE" {
		t.Errorf("formatNameAsEnv(\"PascalCase\") = %v; want PASCAL_CASE", got)
	}
}
