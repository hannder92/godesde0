package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//	fmt.Println(paises)

	//Asignacion por clave - directa
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	//Asignacion en declaracion
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	//Borrado por clave
	delete(campeonato, "Chivas")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Real Madrid"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t\n", puntaje, existe)
}
