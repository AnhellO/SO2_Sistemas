package main

import("bufio"
    "fmt"
    "os"
    "strings")

func main(){
reader := bufio.NewReader(os.Stdin) // para capturar string
fmt.Print("Escriba una Palabra  ") // solo un input
        palabra, _ := reader.ReadString('\n') // capturar string        
        palabra = strings.Replace(palabra, "\n", "", -1)
        palabra = strings.ToLower(palabra)

fmt.Print("Escriba una Vocal  ") // solo un input
        vocal, _ := reader.ReadString('\n') // capturar string
        vocal = strings.Replace(vocal, "\n", "", -1)
       	vocal = strings.ToLower(vocal)

     resultadoe := strings.Replace(palabra,"e", vocal, -1) // cambio vocal por vocal y almaceno
     resultadoi := strings.Replace(resultadoe, "i", vocal, -1)// cambio vocal por vocal y almaceno
     resultadoo := strings.Replace(resultadoi, "o", vocal, -1)// cambio vocal por vocal y almaceno
     resultadou := strings.Replace(resultadoo, "u", vocal, -1)// cambio vocal por vocal y almaceno

 			
        userFile := (resultadou)
		fout, err := os.Create(userFile)		
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		fout.WriteString(resultadou)    

}//cierro main

