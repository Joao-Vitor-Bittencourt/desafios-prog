package desafios

import "fmt"

/*
	Um número é feliz quando segue as seguintes regras:

	- numero começa com qualquer inteiro que seja positivo
	- logo após devemos substituir o número pelo quadrado de cada algarismo: ex: se o número é 89 ficará: 8² 9²
	- repetimos esse processo até que o resultado seja 1, se o resultado for 1 ele resultará em um número feliz.
*/

func HappyNumber(n int) bool {
	switch n {
	case 1, 7:
		return true
	case 0, 2, 3, 4, 5, 6, 8, 9:
		return false
	default:
		return HappyNumber(HappyVerify(n))
	}
}

func HappyVerify(n int) int {
	var out int
	var last int
	for n > 0 {
		last = n % 10
		fmt.Printf("Last: %d \n", last)
		out += last * last
		fmt.Printf("Out: %d \n", out)
		fmt.Printf("N: %d \n", n)
		n /= 10
		fmt.Printf("N final: %d \n", n)
	}
	return out
}
