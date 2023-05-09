package desafios

func Triangulo_Pascal(numeroLinhas int) [][]int {
	res := make([][]int, 0, numeroLinhas)
	if numeroLinhas == 0 {
		return res
	}

	res = append(res, []int{1})
	if numeroLinhas == 1 {
		return res
	}

	res = append(res, []int{1, 1})
	if numeroLinhas == 2 {
		return res
	}

	for i := 2; i < numeroLinhas; i++ {
		top, row := res[len(res)-1], NovaLinha(i+1)
		for j := 1; j < i; j++ {
			row[j] = top[j-1] + top[j]
		}
		res = append(res, row)
	}
	return res
}

func NovaLinha(size int) []int {
	row := make([]int, size)
	row[0], row[size-1] = 1, 1
	return row
}
