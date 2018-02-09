package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Leyendo la entrada...")
    fmt.Println("---------------------")

    for {
        fmt.Print("Ingresa tu texto --> ")
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        
        if strings.Compare("holi", text) == 0 {
            fmt.Println("Hola, humano!")
        } else if strings.Compare("salir", text) == 0 {
            break;
        } else {
            fmt.Println(text)
        }
    }
}