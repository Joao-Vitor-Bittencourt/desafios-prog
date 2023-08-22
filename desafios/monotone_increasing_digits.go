package desafios

import "fmt"

func MonotoneIncreasignDigits(num int) int {
	var valor int
	// agora devemos retornar o maior número possível que seja menor ou igual a num e seja um digito crescente monotono.
	if VerifyMonotone(num) {
		return num
	}

	for i := num; !VerifyMonotone(i + 1); i-- {
		if VerifyMonotone(i) {
			fmt.Print(i)
			valor = i
		}
	}
	return valor
}

func VerifyMonotone(num int) bool {
	// Um inteiro é um digito crescente monótono se e só se cada par de digitos adjacentes x e y validarem x <= y.
	slice := splitInt(num)
	var monotono bool

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] <= slice[i+1] {
			monotono = true
		} else {
			monotono = false
			break
		}

	}
	return monotono
}

func splitInt(n int) []int {
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
