package main

	import (

		"fmt"
		"os"
		"bufio"
		"strings"
		
	)

	func main() {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingrese Palabra:")
		p, _ := reader.ReadString('\n')
		fmt.Println("Ingrese vocal:")
		v, _ := reader.ReadString('\n')

    p= strings.Replace(p,"a", v, -1)
    p= strings.Replace(p,"e", v, -1)
    p= strings.Replace(p,"i", v, -1)
    p= strings.Replace(p,"o", v, -1)
    p= strings.Replace(p,"u", v, -1)

    
    userFile := "archivo.txt"
		fout, err := os.Create(userFile)		
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
			
                fout.WriteString(p)  
   
     
}
