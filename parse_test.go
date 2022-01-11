package envconfig

import "testing"

func TestParseBool(t *testing.T) {
	got, err := parseBool("true")
	if err != nil {
		t.Error(err)
	} else if got != true {
		t.Errorf("parseBool(\"true\") = %v; want true", got)
	}

	got, err = parseBool("false")
	if err != nil {
		t.Error(err)
	} else if got != false {
		t.Errorf("parseBool(\"false\") = %v; want false", got)
	}

	got, err = parseBool("True")
	if err != nil {
		t.Error(err)
	} else if got != true {
		t.Errorf("parseBool(\"True\") = %v; want true", got)
	}

	got, err = parseBool("False")
	if err != nil {
		t.Error(err)
	} else if got != false {
		t.Errorf("parseBool(\"False\") = %v; want false", got)
	}

	got, err = parseBool("1")
	if err != nil {
		t.Error(err)
	} else if got != true {
		t.Errorf("parseBool(\"1\") = %v; want true", got)
	}

	got, err = parseBool("0")
	if err != nil {
		t.Error(err)
	} else if got != false {
		t.Errorf("parseBool(\"0\") = %v; want false", got)
	}

	_, err = parseBool("3")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseBool("abrakadabra")
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestParseInt64(t *testing.T) {
	got, err := parseInt64("0")
	if err != nil {
		t.Error(err)
	} else if got != 0 {
		t.Errorf("parseInt64(\"0\") = %v; want 0", got)
	}

	got, err = parseInt64("-000")
	if err != nil {
		t.Error(err)
	} else if got != 0 {
		t.Errorf("parseInt64(\"-000\") = %v; want 0", got)
	}

	got, err = parseInt64("12345")
	if err != nil {
		t.Error(err)
	} else if got != 12345 {
		t.Errorf("parseInt64(\"12345\") = %v; want 12345", got)
	}

	got, err = parseInt64("+12345")
	if err != nil {
		t.Error(err)
	} else if got != 12345 {
		t.Errorf("parseInt64(\"+12345\") = %v; want 12345", got)
	}

	got, err = parseInt64("-12345")
	if err != nil {
		t.Error(err)
	} else if got != -12345 {
		t.Errorf("parseInt64(\"-12345\") = %v; want -12345", got)
	}

	got, err = parseInt64("9223372036854775807")
	if err != nil {
		t.Error(err)
	} else if got != 9223372036854775807 {
		t.Errorf("parseInt64(\"9223372036854775807\") = %v; want 9223372036854775807", got)
	}

	got, err = parseInt64("-9223372036854775808")
	if err != nil {
		t.Error(err)
	} else if got != -9223372036854775808 {
		t.Errorf("parseInt64(\"-9223372036854775808\") = %v; want -9223372036854775808", got)
	}

	_, err = parseInt64("abrakadabra")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseInt64("12 34")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseInt64("12,34")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseInt64("9223372036854775808")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseInt64("-9223372036854775809")
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestParseUint64(t *testing.T) {
	got, err := parseUint64("0")
	if err != nil {
		t.Error(err)
	} else if got != 0 {
		t.Errorf("parseUint64(\"0\") = %v; want 0", got)
	}

	got, err = parseUint64("000")
	if err != nil {
		t.Error(err)
	} else if got != 0 {
		t.Errorf("parseUint64(\"-000\") = %v; want 0", got)
	}

	got, err = parseUint64("12345")
	if err != nil {
		t.Error(err)
	} else if got != 12345 {
		t.Errorf("parseUint64(\"12345\") = %v; want 12345", got)
	}

	got, err = parseUint64("18446744073709551615")
	if err != nil {
		t.Error(err)
	} else if got != 18446744073709551615 {
		t.Errorf("parseUint64(\"18446744073709551615\") = %v; want 18446744073709551615", got)
	}

	_, err = parseUint64("abrakadabra")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseUint64("12 34")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseUint64("12,34")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseUint64("18446744073709551616")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseUint64("-123")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = parseUint64("+123")
	if err == nil {
		t.Errorf("error expected")
	}
}
