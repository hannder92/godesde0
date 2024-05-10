package files

import (
	"bufio"
	"fmt"
	"github.com/hannder92/godesde0/ejercicios"
	"os"
)

var filename = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Ejercicio2()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Ejercicio2()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al concatenar el contenido en el archivo: ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	//Lee todo el archivo
	/*archivo, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err.Error())
		return
	}
	fmt.Println(string(archivo))
	*/
	//Lee Linea por Linea
	arch, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err.Error())
		return
	}
	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">", registro)
	}
}
