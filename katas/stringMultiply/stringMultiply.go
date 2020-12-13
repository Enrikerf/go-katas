package main

import (
	"math"
	"strings"
)

func main() {
	println(stringMultiply("14", "123"))
}

func stringMultiply(multiplier string, multiplicand string) (result string) {
	if(multiplier[0] == '0' || multiplicand[0]== '0'){
		result = "0"
		return
	}
	var product []int
	product = make([]int, len(multiplier)+len(multiplicand))
	for multiplierPosition := len(multiplier) - 1; multiplierPosition >= 0; multiplierPosition-- {
		for multiplicandPosition := len(multiplicand) - 1; multiplicandPosition >= 0; multiplicandPosition-- {
			positionResult := product[multiplierPosition+multiplicandPosition+1] + runeMultiplier(multiplier[multiplierPosition], multiplicand[multiplicandPosition])
			product[multiplierPosition+multiplicandPosition+1] = int(math.Mod(float64(positionResult), 10))
			product[multiplierPosition+multiplicandPosition] += positionResult / 10
		}
	}
	result = stringBuilder(product)
	if( result[0] == '0'){
		result = result[1:]
	}
	return
}

func runeMultiplier(x byte, y byte) (result int) {
	result = int(x-'0') * int(y-'0')
	return
}

func stringBuilder(array []int) (result string) {
	var stringResult strings.Builder
	for _, value := range array {
		stringResult.WriteRune(rune(value + 48))
	}
	result = stringResult.String()
	return
}
