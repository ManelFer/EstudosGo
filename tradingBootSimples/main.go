package main

import (
	"fmt"
	"trading-bot-simples/services"
	"trading-bot-simples/utils"
)

func main() {
	fmt.Println("Iniciando bot...")

	// Perguntar símbolo
	symbol := utils.AskSymbol()

	// Buscar o preço
	price, err := services.GetPrice(symbol)
	if err != nil {
		fmt.Println("Erro ao buscar preço:", err)
		return
	}

	fmt.Printf("Preço de %s agora: %.2f\n", symbol, price)
}
