package main

func Soma(numeros []int) int {
	soma := 0

	for _, v := range numeros {
		soma += v
	}

	return soma
}

func SomaTudo(slicesParaSomar ...[]int) []int {
	var sliceRetorno []int
	
	for _, v := range slicesParaSomar {
		sliceRetorno = append(sliceRetorno, Soma(v))
	}

	return sliceRetorno
}

func main() {
	
}