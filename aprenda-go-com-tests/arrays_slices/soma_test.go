package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	
	t.Run("Test com slice", func (t *testing.T)  {
		slice := []int{1,2,3}
		resultado := Soma(slice)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, slice)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	t.Run("Test recebendo vários slices", func (t *testing.T)  {
		resultado := SomaTudo([]int{1,2,3}, []int{2,5})
		esperado := []int{6, 7}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
		}
	})
}