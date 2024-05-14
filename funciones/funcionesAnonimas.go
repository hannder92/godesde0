package funciones

import "fmt"

func Calculos() {
	var num3 int = 32
	var num4 int = 243
	calculo := func(a int, b int) int {
		return a + b + num3 + num4
	}
	fmt.Println(calculo(1, 2))

	calculo = func(a int, b int) int {
		return a * b * num3
	}
	fmt.Println(calculo(1, 2))
}
