package main

import "testing"

func TestSoma(t *testing.T){
	total := Somar(2,5)

	if total != 7 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d",total,30)
	}
}

func TestSub(t *testing.T){
	total := Sub(5,2)

	if total != 3 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d",total,30)
	}
}

func TestTimes(t *testing.T){
	total := Times(10,5)

	if total != 2 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: %d",total,30)
	}
}

func TestSumX(t *testing.T){
	total := SumX(2,5)

	if total != 10 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d",total,30)
	}
}