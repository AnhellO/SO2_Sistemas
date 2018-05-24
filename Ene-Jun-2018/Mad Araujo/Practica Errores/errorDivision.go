package main

import "fmt"
import "errors"
import "strings"
import "strconv"
import "bufio"
import "os"

func division(valor1 int, valor2 int) (int, error) {
    if valor1 == 0 || valor2 == 0 {
        return 1, errors.New("No se puede dividir por 0")
    }

    return valor1 / valor2, nil;
}

type errorCustom struct {
    valor1, valor2 int
    mensaje string
}


func main() {
    reader := bufio.NewReader(os.Stdin) // para capturar string

    // Probando con la interface Errors de Go
    fmt.Print("Ingrese el Divisior\n") // solo un input
        texta, _ := reader.ReadString('\n') // capturar string
        texta = strings.Replace(texta, "\n", "", -1) // capturar string
        a, _ := strconv.Atoi(texta) /// convertimos el string a entero

        fmt.Print("Ingrese el Dividendo\n") // solo un input
        textb, _ := reader.ReadString('\n') // capturar string
        textb = strings.Replace(textb, "\n", "", -1) // capturar string
        b, _ := strconv.Atoi(textb) /// convertimos el string a entero
   
    resultado, err := division(a, b)
    fmt.Printf("Resultado = %d\n", resultado)

    if err != nil {
        fmt.Println(err)
    }

    
}