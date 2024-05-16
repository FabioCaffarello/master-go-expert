package exchangerateentity

import (
	"fmt"
	"strconv"
)

// StringToFloat64 converts a string to a float64.
// It returns the float64 value and an error if the conversion fails.
func StringToFloat64(s string) (float64, error) {
	// Use strconv.ParseFloat to convert the string to a float64
	// The second parameter, 64, specifies that we want a float64
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to float64: %w", err)
	}
	return value, nil
}


// StringToInt64 converts a string to an int64.
// It returns the int64 value and an error if the conversion fails.
func StringToInt64(s string) (int64, error) {
    // Use strconv.ParseInt to convert the string to an int64
    // The second parameter, 10, specifies that the number is in base 10
    // The third parameter, 64, specifies that we want an int64
    value, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return 0, fmt.Errorf("failed to convert string to int64: %w", err)
    }
    return value, nil
}
