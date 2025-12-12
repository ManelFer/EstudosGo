package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// GetPrice busca o preço na Binance
func GetPrice(symbol string) (float64, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data BinanceResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	// Se a API retornar erro (símbolo inválido)
	if data.Price == "" {
		return 0, fmt.Errorf("símbolo inválido ou não encontrado")
	}

	// Converter string em float64
	return strconv.ParseFloat(data.Price, 64)
}
