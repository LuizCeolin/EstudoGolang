package main

import (
	"testing"
)

func TestAdicionar(t *testing.T) {

	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Soma '%d', esperado '%d'", soma, esperado)
	}

}