package number

import "math"

// Round 浮点数四舍五入
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}
