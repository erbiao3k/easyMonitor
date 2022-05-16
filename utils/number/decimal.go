package number

import (
	"fmt"
	"strconv"
)

// Decimal 浮点数精度调整为2
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
