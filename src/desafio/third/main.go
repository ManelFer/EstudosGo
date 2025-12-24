// criar um programa em go que trabalhe com dois pacotes e concorrencia.
// escrever pin pong
// ultilizar goroutines e channels exiba em alternancia pin pong

package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	for {
		ch <- "ping"
		time.Sleep(1 * time.Second)
	}
}

func pong(ch chan string) {
	for {
		ch <- "pong"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan string)

	// goroutine recebe e exibe as mensagens do channel
	go ping(channel)
	go pong(channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}
