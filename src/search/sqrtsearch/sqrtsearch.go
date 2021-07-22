package sqrtsearch

import (
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

func Sqrtsearch(target decimal.Decimal) float64 {
	left, right := decimal.NewFromInt(0), target
	result := 0.1
	absResult := 0.1
	for left.LessThan(right) {
		mid := right.Add(left).Div(decimal.NewFromInt(2))
		absResult, _ = mid.Mul(mid).Sub(target).Float64()
		if math.Abs(absResult) <= 0.000001 {
			result, _ = mid.Float64()
			break
		}
		if mid.Mul(mid).Equal(target) {
			f, _ := mid.Float64()
			result = f
			break
		} else if mid.Mul(mid).GreaterThan(target) {
			if mid == right {
				f, _ := mid.Float64()
				result = f
				break
			}
			right = mid
		} else {
			left = mid
		}
	}
	return result
}

func PrecisionCheck(expectPrecision int, target decimal.Decimal) bool{
	str := target.String()
	strArr := strings.Split(str, ".")
	if len(strArr[1]) == expectPrecision {
		return true
	}
	return false
}

/*
2.5 1.25 1.875 2.1875 2.34375 2.265625 2.2265625 2.24609375 2.236328125
n/2 n/4 n/8 .... n/2^k
*/
