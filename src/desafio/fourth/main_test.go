// arquivo teste
package main

import (
	"testing"
)

func TestLerArquivoJSON(t *testing.T) {
	data, err := lerAquivoJSON()
	if err != nil {
		t.Fatalf("Erro ao ler arquivo JSON: %v", err)
	}
	if len(data.Clientes) == 0 {
		t.Error("Esperava-se que o arquivo JSON contenha clientes")
	}
}
