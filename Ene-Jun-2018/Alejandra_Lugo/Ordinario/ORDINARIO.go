package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	//valid "github.com/asaskevich/govalidator"
)

func main() {

	var opcion, a, b, resultado float32

	fmt.Println("Teclea A")
	fmt.Scan("%g", &a)
	fmt.Println("Teclea B")
	fmt.Scan("%g", &b)

	if reflect.TypeOf(a) = int {
		if reflect.TypeOf(b)= int{
fmt.Println("Elige la Opcion")
	fmt.Println("1. Sumar")
	fmt.Println("2. Restar")
	fmt.Println("3. Multiplicar")
	fmt.Println("4. Dividir")
	fmt.Println("5. Raíz")
	fmt.Println("6. Potencia")
	fmt.Println("7. Salir")
	fmt.Scan(&opcion)

	switch opcion{
	case 1:
		resultado=a+b
		fmt.Println("Resultado = ", resultado)
	case 2:
		resultado=a-b
		fmt.Println("Resultado = ", resultado)
	case 3:
    	resultado=a*b
    	fmt.Println("Resultado = ", resultado)
    case 4:
	if b == 0 {
			err := errors.New("No te pases, B es 0!")
			if err != nil {
				fmt.Println(err)
			}
		}
		if b != 0 {
			resultado = a / b
			fmt.Println("Resultado =", resultado)
		}
	case 5:
		resultado = math.Pow(a,1/b)
		fmt.Println("Resultado = ", resultado)
	case 6:
		resultado = math.Pow(a,b)
		fmt.Println("Resultado = ", resultado)
	case 7:
		break	
	}		

		}
		
	} else {
	err := errors.New("El valor no es correcto!")
			if err != nil {
				fmt.Println(err)
			}
	}
}