package main

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade uint
}

func main() {

	var p = Pessoa{}

	fmt.Println("Qual é o seu nome?")
	fmt.Scanf("%s\n", &p.nome)

	fmt.Println("Qual é a sua idade?")
	fmt.Scanf("%d\n", &p.idade)
	fmt.Printf("ola %s, sua idade é %d\n", p.nome, p.idade)
}
