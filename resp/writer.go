package resp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	StringType string = "string"
	NumberType string = "number"
	ErrorType  string = "error"
	ArrayType  string = "array"
)

func ParseRESPValue(value string) (string, error) {
	valueType := getValueType(value)
	switch valueType {
	case NumberType:
		inValue, err := strconv.Atoi(value)
		if err == nil {
			return EncodeRESPNumber(int64(inValue)), nil
		}
	case StringType:
		return EncodeRESPString(value), nil
	case ArrayType:
		return EncodeRESPArray(value), nil
	}
	fmt.Println(value, valueType)
	return "", nil
}

// EncodeRESPNumber encodes a given integer as a RESP integer.
func EncodeRESPNumber(Number int64) string {
	return fmt.Sprintf(":%d\r\n", Number)
}

// EncodeRESPNumber encodes a given integer as a RESP integer.
func EncodeRESPString(str string) string {
	if len(str) > 2 {
		// Bulk string.
		return fmt.Sprintf("$%d\r\n%s\r\n", len(str), str)
	} else {
		// Simple string
		return fmt.Sprintf("+%s\r\n", str)
	}
}

func EncodeRESPArray(s string) string {
	// Remove the square brackets from the input string.
	s = strings.Trim(s, "[]")

	// Split the string into individual values.
	values := strings.Split(s, ",")

	// Parse each value and add it to the result array.
	var result []interface{}
	for _, v := range values {
		v = strings.TrimSpace(v)
		decoded, err := ParseRESPValue(v)
		if err != nil {

		}
		result = append(result, decoded)
	}
	if len(result) > 0 {
		return fmt.Sprintf("+%d\r\n%s\r\n", len(result), result)
	} else {
		return fmt.Sprintf("+%d\r\n%s\r\n", len(result), "")
	}
}

// Switch on the value and get it's type, all types should be encoded in a string type.
func getValueType(value string) string {
	switch value[0] {
	case '{':
		return "hash" // TODO: This type is not required for now.
	case '[':
		return ArrayType
	default:
		_, err := strconv.Atoi(value)
		if err == nil {
			return NumberType
		} else {
			return StringType
		}
	}
}

// Check the value of the string and
func checkStringType(str string) string {
	length := len(str)
	if length > 2 {
		return "bulk string"
	}
	return "simple string"
}
