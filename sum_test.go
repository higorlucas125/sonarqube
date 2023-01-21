package main

import "testing"

func TestSoma(t *testing.T){
	total := Somar(2,5)

	if total != 7 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d",total,30)
	}
}