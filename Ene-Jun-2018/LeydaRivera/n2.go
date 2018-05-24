package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	//w := bufio.NewWriter(os.Stdout)
	fmt.Println("Ingresa un nùmero:")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		userFile := "text" + strconv.Itoa(i) + ".txt" //para que el nombre del archivo vaya cambiando segun i
		fout, err := os.Create(userFile)
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		fout.WriteString("HOLA\r\n")
	}

}
