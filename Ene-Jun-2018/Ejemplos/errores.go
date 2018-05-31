package main

import "fmt"
import "errors"

func sumaNegativa(valor1 int, valor2 int) (int, error) {
    if valor1 >= 0 || valor2 >= 0 {
        return 1, errors.New("Te estás pasando de positivo!")
    }

    return valor1 + valor2, nil;
}

type errorCustom struct {
    valor1, valor2 int
    mensaje string
}

func (e *errorCustom) Error() string {
    return fmt.Sprintf("Error:\nValor A: %d\nValor B:%d\nMensaje: %s", e.valor1, e.valor2, e.mensaje)
}

func sumaPositiva(valor1 int, valor2 int) (int, error) {
    if valor1 < 0 || valor2 < 0 {
        return -1, &errorCustom{valor1, valor2, "Te estás pasando de negativo!"}
    }

    return valor1 + valor2, nil;
}

func main() {
    // Probando con la interface Errors de Go
    a := -88
    b := -5
    resultadoPositivo, err := sumaNegativa(a, b)
    fmt.Printf("Resultado = %d\n", resultadoPositivo)

    if err != nil {
        fmt.Println(err)
    }

    // Probando con structs
    c := 10
    d := -5
    resultadoNegativo, err2 := sumaPositiva(c, d)
    fmt.Printf("Resultado = %d\n", resultadoNegativo)

    if err2 != nil {
        fmt.Println(err2)
    }
}
