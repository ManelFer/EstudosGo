package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AskSymbol lê do terminal e retorna o nome da cripto
func AskSymbol() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o símbolo da cripto (ex: BTCUSDT): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return strings.ToUpper(input)
}
