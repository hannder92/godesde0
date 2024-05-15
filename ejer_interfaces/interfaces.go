package ejer_interfaces

import (
	"fmt"
	"github.com/hannder92/godesde0/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", humano.Sexo())
}
