package envconfig

import "strconv"

func parseBool(str string) (bool, error) {
	result, err := strconv.ParseBool(str)
	return result, err
}

func parseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func parseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func parseFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
