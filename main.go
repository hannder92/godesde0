package main

import (
	"fmt"
	"github.com/hannder92/godesde0/arreglos_slices"
	"github.com/hannder92/godesde0/defer_panic"
	"github.com/hannder92/godesde0/ejer_interfaces"
	"github.com/hannder92/godesde0/ejercicios"
	"github.com/hannder92/godesde0/funciones"
	"github.com/hannder92/godesde0/iteraciones"
	"github.com/hannder92/godesde0/mapas"
	"github.com/hannder92/godesde0/middleware"
	"github.com/hannder92/godesde0/modelos"
	"github.com/hannder92/godesde0/users"
	"github.com/hannder92/godesde0/variables"
	"runtime"
)

func main() {
	//Variables
	variables.MostrarEnteros()
	variables.RestoVariables()

	//Funciones
	estado, texto := variables.ConvertirATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	//Condicionales
	//IF
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es:", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	//SWITCH
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Esto es darwin")
	case "linux":
		fmt.Println("Esto es linux")
	default:
		fmt.Printf("%s\n", os)
	}

	entero, respuesta := ejercicios.Ejercicio1("1000")
	fmt.Println(entero)
	fmt.Println(respuesta)

	//Ingreso por teclado - bufio
	//teclado.IngresoNumeros()

	//Iteraciones - FOR
	iteraciones.Iterar()

	//Ejercicio Iteracion
	//fmt.Println(ejercicios.Ejercicio2())

	//Escritura en archivos
	//files.GrabaTabla()
	//files.SumaTabla()
	//Lectura de Archivo
	//files.LeoArchivo()

	//Funciones anonimas
	funciones.Calculos()

	//Closure
	funciones.LlamarClosure()

	//Recursion
	funciones.Exponencia(2)

	//Arreglos
	arreglos_slices.MuestroArreglos()

	//Slices
	arreglos_slices.MuestroSlice()

	arreglos_slices.Capacidad()

	//Mapas
	mapas.MostrarMapas()

	//Estructura
	users.AltaUsuario()

	//Interfaces
	Pedro := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Maria)

	//Defer
	defer_panic.VemosDefer()

	//Panic
	//defer_panic.EjemploPanic()

	//Asincronia
	/*canal1 := make(chan bool)
	go goroutines.MiNombreLento("Pepito Perez", canal1)
	//si hay mas de un canal, se pueden poner en el defer para que espere todas las funciones
	defer func() {
		<-canal1
	}()
	//<-canal1*/

	//Servidor Web
	//webserver.MiWebServer()

	middleware.MiMiddleware()

}
