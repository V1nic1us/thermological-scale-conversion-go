package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Digite o valor da temperatura em Fahrenheit: ")
	fmt.Scan(&input)

	fahrenheit, err := strconv.ParseInt(input, 10, 64)
    if err != nil {
        fmt.Println("Erro ao converter entrada para número:", err)
        return
    }
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("Temperatura em graus Celsius",celsius)
}