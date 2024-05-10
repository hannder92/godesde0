package iteraciones

import "fmt"

func Iterar() {

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
		//continue - continua la siguiente iteracion y salta el resto del flujo
		//break - sale del for
	}
}
