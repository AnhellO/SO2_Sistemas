package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var v, pa, resultado string
	fmt.Println("Ingresa Palabra:")
	fmt.Scan(&pa)
	fmt.Println("Ingresa vocal:")
	fmt.Scan(&v)

	//remplazar
	if v == "a" {
		cambio := strings.NewReplacer("a", "a", "e", "a", "i", "a", "o", "a", "u", "a")
		resultado = cambio.Replace(pa)
		fmt.Println(resultado)
	}

	if v == "e" {
		cambio := strings.NewReplacer("a", "e", "e", "e", "i", "e", "o", "e", "u", "e")
		resultado = cambio.Replace(pa)
		fmt.Println(resultado)
	}

	if v == "i" {
		cambio := strings.NewReplacer("a", "i", "e", "i", "i", "i", "o", "i", "u", "i")
		resultado = cambio.Replace(pa)
		fmt.Println(resultado)
	}

	if v == "o" {
		cambio := strings.NewReplacer("a", "o", "e", "o", "i", "o", "o", "o", "u", "o")
		resultado = cambio.Replace(pa)
		fmt.Println(resultado)
	}

	if v == "u" {
		cambio := strings.NewReplacer("a", "u", "e", "u", "i", "u", "o", "u", "u", "u")
		resultado = cambio.Replace(pa)
		fmt.Println(resultado)
	}

	userFile := resultado + strconv.Itoa(1) + ".txt" //para que el nombre del archivo vaya cambiando segun i
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	fout.WriteString(resultado)
}
