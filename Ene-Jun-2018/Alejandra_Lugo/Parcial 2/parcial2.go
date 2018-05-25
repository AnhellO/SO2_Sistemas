package main

import (
	"errors"
	"fmt"
)

func main() {

	//var comida string
	var comida, opcion, cantidad, Hamburguesa, Hot, Tacos int

	Hamburguesa = 10
	Hot = 5
	Tacos = 15
	//	opcion := 4

	fmt.Println("Elige la Opcion")
	fmt.Println("Opcion 1: Pedir Comida")
	fmt.Println("Opcion 2: Surtir")
	fmt.Println("Opcion 3: Imprimir Comida")
	fmt.Println("Opcion 4: Salir")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		fmt.Println("¿Qué Comida eliges?")
		fmt.Println("1. Hamburguesa")
		fmt.Println("2. Hot dog")
		fmt.Println("3. Tacos")
		fmt.Scan(&comida)

		switch comida {
		case 1:
			if Hamburguesa == 0 {
				err := errors.New("Ya no tenemos Hamburguesas, una disculpa.")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Disfruta tu Hamburguesa")
			}
			Hamburguesa--

		case 2:
			if Hot == 0 {
				err := errors.New("Ya no tenemos Hot dogs, una disculpa.")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Disfruta tu Hot dog")
			}

			Hot--

		case 3:
			if Tacos == 0 {
				err := errors.New("Ya no tenemos Tacos, una disculpa.")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Disfruta tus Tacos")
			}

			Tacos--
		}

	case 2:
		fmt.Println("¿Qué Comida vas a surtir?")
		fmt.Println("1. Hamburguesa")
		fmt.Println("2. Hot dog")
		fmt.Println("3. Tacos")
		fmt.Scan(&comida)
		fmt.Println("Cantidad:")
		fmt.Scan(&cantidad)

		if cantidad <= 1 {
			err := errors.New("Falló al surtir la comida")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			switch comida {
			case 1:
				Hamburguesa += cantidad
			case 2:
				Hot += cantidad
			case 3:
				Tacos += cantidad
			}
			fmt.Println("Comida surtida con éxito :D")
		}

	case 3:
		fmt.Println("Hamburguesas: ", Hamburguesa)
		fmt.Println("Hot dog: ", Hot)
		fmt.Println("Tacos: ", Tacos)

	case 4:
		break

	default:
		fmt.Println("No elegiste una opción válida :V")
	}

}
