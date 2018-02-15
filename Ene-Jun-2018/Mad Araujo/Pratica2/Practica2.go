package main

import("bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "io/ioutil")

func main(){
reader := bufio.NewReader(os.Stdin) // para capturar string
fmt.Print("Cantidad de Archivos que desea crear de 1 al 5  ") // solo un input
        text, _ := reader.ReadString('\n') // capturar string
        text = strings.Replace(text, "\n", "", -1) // capturar string
       z_entero, _ := strconv.Atoi(text) /// convertimos el string a entero
        if z_entero == 1 {						//condicional para crear 1 archivo
        //for i := 0; i <= z_entero; i++{ 	
         	userFile := "Archivo.txt"
		fout, err := os.Create(userFile)		
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
			fout.WriteString("Escogiste solo crear 1\r\n") 
			//lee archivo
			b, err := ioutil.ReadFile("Archivo.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			str := string(b) // convierto de byts a string
    			fmt.Println(str) // imprimo lo leido
         }

         if z_entero == 2 {				//condicional para crear 2 archivo
        //for i := 0; i <= z_entero; i++{ 	
         	userFile1 := "Archivo1.txt"
		fout, err := os.Create(userFile1)		
		if err != nil {
			fmt.Println(userFile1, err)
			return
		}
		defer fout.Close()
			fout.WriteString("Este es el Pirmer Archivo\r\n") 
			//lee archivo
			b, err := ioutil.ReadFile("Archivo1.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			str := string(b) // convierto de byts a string
    			fmt.Println(str) // imprimo lo leido
         //creamos archivo2
         userFile2 := "Archivo2.txt"
		fout2, err := os.Create(userFile2)		
		if err != nil {
			fmt.Println(userFile2, err)
			return
		}
		defer fout2.Close()
			fout2.WriteString("Este es el Segundo Archivo\r\n")
			//lee archivo
			c, err := ioutil.ReadFile("Archivo2.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strc := string(c) // convierto de byts a string
    			fmt.Println(strc) // imprimo lo leido

         }///cierre if 2

         if z_entero == 3 {				//condicional para crear 3 archivo
        //for i := 0; i <= z_entero; i++{ 	
         	userFile1 := "Archivo1.txt"
		fout, err := os.Create(userFile1)		
		if err != nil {
			fmt.Println(userFile1, err)
			return
		}
		defer fout.Close()
			fout.WriteString("Este es el Pirmer Archivo\r\n") 
			//lee archivo
			b, err := ioutil.ReadFile("Archivo1.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			str := string(b) // convierto de byts a string
    			fmt.Println(str) // imprimo lo leido
         //creamos archivo2
         userFile2 := "Archivo2.txt"
		fout2, err := os.Create(userFile2)		
		if err != nil {
			fmt.Println(userFile2, err)
			return
		}
		defer fout2.Close()
			fout2.WriteString("Este es el Segundo Archivo\r\n")
			//lee archivo
			c, err := ioutil.ReadFile("Archivo2.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strc := string(c) // convierto de byts a string
    			fmt.Println(strc) // imprimo lo leido
			//creamos archivo3
         userFile3 := "Archivo3.txt"
		fout3, err := os.Create(userFile3)		
		if err != nil {
			fmt.Println(userFile3, err)
			return
		}
		defer fout3.Close()
			fout3.WriteString("Este es el Tercer Archivo\r\n")
			//lee archivo
			d, err := ioutil.ReadFile("Archivo3.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strd := string(d) // convierto de byts a string
    			fmt.Println(strd) // imprimo lo leido

         }///cierre if 3

         if z_entero == 4 {				//condicional para crear 4 archivo
        //for i := 0; i <= z_entero; i++{ 	
         	userFile1 := "Archivo1.txt"
		fout, err := os.Create(userFile1)		
		if err != nil {
			fmt.Println(userFile1, err)
			return
		}
		defer fout.Close()
			fout.WriteString("Este es el Pirmer Archivo\r\n") 
			//lee archivo
			b, err := ioutil.ReadFile("Archivo1.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			str := string(b) // convierto de byts a string
    			fmt.Println(str) // imprimo lo leido
         //creamos archivo2
         userFile2 := "Archivo2.txt"
		fout2, err := os.Create(userFile2)		
		if err != nil {
			fmt.Println(userFile2, err)
			return
		}
		defer fout2.Close()
			fout2.WriteString("Este es el Segundo Archivo\r\n")
			//lee archivo
			c, err := ioutil.ReadFile("Archivo2.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strc := string(c) // convierto de byts a string
    			fmt.Println(strc) // imprimo lo leido
			//creamos archivo3
         userFile3 := "Archivo3.txt"
		fout3, err := os.Create(userFile3)		
		if err != nil {
			fmt.Println(userFile3, err)
			return
		}
		defer fout3.Close()
			fout3.WriteString("Este es el Tercer Archivo\r\n")
			//lee archivo
			d, err := ioutil.ReadFile("Archivo3.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strd := string(d) // convierto de byts a string
    			fmt.Println(strd) // imprimo lo leido

			//creamos archivo4
         userFile4 := "Archivo4.txt"
		fout4, err := os.Create(userFile4)		
		if err != nil {
			fmt.Println(userFile4, err)
			return
		}
		defer fout4.Close()
			fout4.WriteString("Este es el Cuarto Archivo\r\n")
			//lee archivo
			e, err := ioutil.ReadFile("Archivo4.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			stre := string(e) // convierto de byts a string
    			fmt.Println(stre) // imprimo lo leido

         }///cierre if 4

         if z_entero == 5 {				//condicional para crear 5 archivo
        //for i := 0; i <= z_entero; i++{ 	
         	userFile1 := "Archivo1.txt"
		fout, err := os.Create(userFile1)		
		if err != nil {
			fmt.Println(userFile1, err)
			return
		}
		defer fout.Close()
			fout.WriteString("Este es el Pirmer Archivo\r\n") 
			//lee archivo
			b, err := ioutil.ReadFile("Archivo1.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			str := string(b) // convierto de byts a string
    			fmt.Println(str) // imprimo lo leido
         //creamos archivo2
         userFile2 := "Archivo2.txt"
		fout2, err := os.Create(userFile2)		
		if err != nil {
			fmt.Println(userFile2, err)
			return
		}
		defer fout2.Close()
			fout2.WriteString("Este es el Segundo Archivo\r\n")
			//lee archivo
			c, err := ioutil.ReadFile("Archivo2.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strc := string(c) // convierto de byts a string
    			fmt.Println(strc) // imprimo lo leido
			//creamos archivo3
         userFile3 := "Archivo3.txt"
		fout3, err := os.Create(userFile3)		
		if err != nil {
			fmt.Println(userFile3, err)
			return
		}
		defer fout3.Close()
			fout3.WriteString("Este es el Tercer Archivo\r\n")
			d, err := ioutil.ReadFile("Archivo3.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strd := string(d) // convierto de byts a string
    			fmt.Println(strd) // imprimo lo leido

			//creamos archivo4
         userFile4 := "Archivo4.txt"
		fout4, err := os.Create(userFile4)		
		if err != nil {
			fmt.Println(userFile4, err)
			return
		}
		defer fout4.Close()
			fout4.WriteString("Este es el Cuarto Archivo\r\n")
			//lee archivo
			e, err := ioutil.ReadFile("Archivo4.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			stre := string(e) // convierto de byts a string
    			fmt.Println(stre) // imprimo lo leido

         //creamos archivo5
         userFile5 := "Archivo5.txt"
		fout5, err := os.Create(userFile5)		
		if err != nil {
			fmt.Println(userFile5, err)
			return
		}
		defer fout5.Close()
			fout5.WriteString("Este es el Quinto Archivo\r\n")
			//lee archivo
			f, err := ioutil.ReadFile("Archivo5.txt") // Paso nombre de archivo
    			if err != nil {
        			fmt.Print(err)
    			}
    			strf := string(f) // convierto de byts a string
    			fmt.Println(strf) // imprimo lo leido


         }///cierre if 5

         

}//cierro main

