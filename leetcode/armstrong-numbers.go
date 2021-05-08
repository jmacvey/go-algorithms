package leetcode

import (
	"math"
)


func IsArmstrong(n int) bool {
    numDigits := getNumDigits(n)
    numDigitsFloat64 := float64(numDigits)
    
    sum := float64(0)
	value := n
    for i := 0; i < numDigits; i++ {
        digit := value % 10
		value /= 10
        sum += math.Pow(float64(digit), numDigitsFloat64)
    }
    
    return sum == float64(n)
}

func getNumDigits(n int) int {
    i := 0
    for n > 0 {
        n /= 10
        i++
    }
	
    
    return i
}