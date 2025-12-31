package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type TodosUsuarios struct {
	TodosUsuarios []Usuarios `json:"Usuarios"`
}

type Usuarios struct {
	Nome   string `json:"Nome"`
	Idade  int    `json:"Idade"`
	Tipo   string `json:"Tipo"`
	Social Social `json:"Social"`
}
type Social struct {
	Facebook  string `json:"Facebook"`
	Twitter   string `json:"Twitter"`
	Linkedin  string `json:"Linkedin"`
	Instagram string `json:"Instagram"`
}

func main() {

	jsonFile, err := os.Open("usuarios.json")

	// se houve um erro ao abrir o arquivo
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var Usuarios TodosUsuarios
	json.Unmarshal(byteValue, &Usuarios)

	for i := 0; i < len(Usuarios.TodosUsuarios); i++ {
		fmt.Println("Nome: " + Usuarios.TodosUsuarios[i].Nome)
		fmt.Println("Idade: " + strconv.Itoa(Usuarios.TodosUsuarios[i].Idade))
		fmt.Println("Tipo: " + Usuarios.TodosUsuarios[i].Tipo)
		fmt.Println("Facebook: " + Usuarios.TodosUsuarios[i].Social.Facebook)
		fmt.Println("Twitter: " + Usuarios.TodosUsuarios[i].Social.Twitter)
		fmt.Println("Linkedin: " + Usuarios.TodosUsuarios[i].Social.Linkedin)
		fmt.Println("Instagram: " + Usuarios.TodosUsuarios[i].Social.Instagram)
		fmt.Println("----------------------------")
	}
}
