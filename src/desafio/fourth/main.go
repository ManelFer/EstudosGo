// O desafio é criar uma API que a vovó possa guardar dados dos seus clientes.
// criar o JSON com as informações dos clientes - feito

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type todosClientes struct {
	Clientes []clientes `json:"clientes"`
}

type clientes struct {
	Nome     string   `json:"nome"`
	Idade    int      `json:"idade"`
	Endereco endereco `json:"endereco"`
}

type endereco struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Bairro string `json:"bairro"`
}

// nesse desafio precisa de um funcao que leia o arquivo JSON
func lerAquivoJSON() (todosClientes, error) {
	arquivo, err := os.Open("./data/client.json")
	if err != nil {
		return todosClientes{}, err
	}
	defer arquivo.Close()

	var data todosClientes
	err = json.NewDecoder(arquivo).Decode(&data)
	if err != nil {
		return todosClientes{}, err
	}
	return data, nil
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	data, err := lerAquivoJSON()
	if err != nil {
		http.Error(w, "Erro ao ler arquivo JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/clientes", getClientes).Methods("GET")
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", route))

}
