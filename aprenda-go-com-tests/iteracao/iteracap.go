package main

const qtdRepeticoes = 5

func Repetir(caracter string) string{
	
	var repeticoes string

	for i := 0; i < qtdRepeticoes; i++ {
		repeticoes += caracter  
	}

	return repeticoes
}

func main() {
	
}