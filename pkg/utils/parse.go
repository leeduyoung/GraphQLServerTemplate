package utils

import "strconv"

// ParseStringToInt64 convert string -> int64
func ParseStringToInt64(data string) int64 {
	parseData, _ := strconv.ParseInt(data, 10, 64)
	return parseData
}

// ParseStringToInt convert string -> int
func ParseStringToInt(data string) int {
	parseData, _ := strconv.Atoi(data)
	return parseData
}

// ParseFloat64ToString convert float64 -> string
func ParseFloat64ToString(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 32)
}

// ParseIntToString convert int -> string
func ParseIntToString(data int) string {
	return strconv.Itoa(data)
}

// ParseInt64ToString convert int64 -> string
func ParseInt64ToString(data int64) string {
	return strconv.FormatInt(data, 10)
}
