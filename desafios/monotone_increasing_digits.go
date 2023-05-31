package desafios

import "fmt"

func MonotoneIncreasignDigits(num int) {

}

func verifyMonotone(num int) bool {
// Um inteiro é um digito crescente monótono se e só se cada par de digitos adjacentes x e y validarem x <= y.
}

func splitInt(num int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}

	return result
}