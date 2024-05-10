package ejercicios

import "strconv"

func Ejercicio1(valor string) (int, string) {
	valorEntero, _ := strconv.Atoi(valor)

	if valorEntero > 100 {
		return valorEntero, "Es mayor a 100"
	} else {
		return valorEntero, "Es menor a 100"
	}

}
