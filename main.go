package main

import (
	"fmt"
	"github.com/hannder92/godesde0/variables"
)

func main() {
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvertirATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
