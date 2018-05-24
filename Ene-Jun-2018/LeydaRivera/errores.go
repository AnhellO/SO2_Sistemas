package main

import (
	"errors"
	"fmt"
	// "os"
	// "strconv"
	//"strings"
)

func main() {
	var n, m, resultado int
	for {
		fmt.Print("N:")
		fmt.Scan(&n)
		fmt.Print("M:")
		fmt.Scan(&m)
		if m == 0 {
			err := errors.New("No te pases,M es 0! :V")
			//fmt.Println("No te pases,M es 0! :V")
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("El resultado es:", resultado)

		}
		if m != 0 {
			resultado = n / m
			fmt.Println("El resultado es:", resultado)

		}

	}
	//break
}
