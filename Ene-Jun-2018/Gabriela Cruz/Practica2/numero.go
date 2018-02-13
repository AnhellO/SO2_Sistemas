package main

	import (
		"fmt"
		"os"
		"bufio"
		"strings"
		"strconv"
		"io/ioutil"
	)

	func main() {


	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingrese Nuemero:")
    fmt.Println("****************")
    
    
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n","", -1)

    
    x, _ := strconv.Atoi(text)

   

    for i:= 0; i<x; i++{

        userFile := "Archivo.txt"
		fout, err := os.Create(userFile)		
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
			
                fout.WriteString("Hola Clase de SO2!\r\n")           
	

	datos,err := ioutil.ReadFile("./Archivo.txt")
	if err !=nil{
		fmt.Println("ERROR")
	}
	fmt.Println(string(datos))

}

    







	


  
}

