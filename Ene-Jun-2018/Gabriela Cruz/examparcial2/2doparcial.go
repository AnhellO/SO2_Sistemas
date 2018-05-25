package main

import("bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "errors")

func surtido(valor1 int, valor2 int) (int, error) {
    if valor1 == 0 || valor2 <= 0  {
        return 0, errors.New("Fallo al surtir Comida")
    }

    return valor1 + valor2, nil;
}

func main(){

	pizza:=5
	jochos:=5
	hambuergueza:=5
	burrito:=5
	var surtidooP, surtidooJ, surtidooH, surtidooB int
	var cantidadP, cantidadJ, cantidadH, cantidadB int

reader := bufio.NewReader(os.Stdin) // para capturar string
for{
fmt.Print("\nRESTAURANTE GABYS\n1.Menu\n2.Surtir Comida\n3.Imprimir comida Disponible\n4.Salir\nElija el Numero de Opcion que desee\n\n") // Opcion a elegir
        text, _ := reader.ReadString('\n') // capturar string
        text = strings.Replace(text, "\n", "", -1) // capturar string
       z_entero, _ := strconv.Atoi(text) /// convertimos el string a entero
        
        if z_entero == 1 {	///Mlegir comida///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nQue comida desea\n1.Pizza\n2.Jochos\n3.Hambuergueza\n4.Burrito\nElija el Numero de Opcion que deseen\n")
      	comida, _ := reader.ReadString('\n') // capturar string
        comida = strings.Replace(comida, "\n", "", -1) // capturar string
        comida = strings.ToLower(comida)
        comidaa, _ := strconv.Atoi(comida) /// convertimos el string a entero
        if comidaa==1 { // condicional pedido
        	fmt.Print("\n¿Cuantas Pizzas Deseas?\n")
        	fmt.Scanf("%d\n", &cantidadP)
        	pizza=pizza-cantidadP
        	if pizza>0{
        		fmt.Print("Hay Pizza!, Existencias = ", pizza,"\n")
        	}else{
        		fmt.Print("No Hay Pizza\n")
        	}
        }

        if comidaa==2 {
        	fmt.Print("\n¿Cuantos Jochos Deseas?\n")
        	fmt.Scanf("%d\n", &cantidadJ)
        	jochos=jochos-cantidadJ
        	if jochos>0{
        		fmt.Print("Hay jochos, Existencias = ", jochos,"\n")
        	}else{
        		fmt.Print("No Hay jochos\n")
        	}
        }

        if comidaa==3 {
        	fmt.Print("\n¿Cuantas Hambuerguezas Deseas?\n")
        	fmt.Scanf("%d\n", &cantidadH)
        	hambuergueza=hambuergueza-cantidadH
        	if hambuergueza>0{
        		fmt.Print("Hay hambuergueza, Existencias = ", hambuergueza,"\n")
        	}else{
        		fmt.Print("No Hay hambuergueza\n")
        	}
        }
        if comidaa==4 {
        	fmt.Print("\n¿Cuantos Burritos Deseas?\n")
        	fmt.Scanf("%d\n", &cantidadB)
        	burrito=burrito-cantidadB
        	if burrito>0{
        		fmt.Print("Hay burrito, Existencias = ", burrito,"\n")
        	}else{
        		fmt.Print("No Hay burrito\n")
        	}
        }
        
    }// cierre Pedido MENU
    /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    if z_entero ==2{ /// COndicional Surtir
    	fmt.Print("\nQue comida desea Surtir\n1.Pizza\n2.Jochos\n3.Hambuergueza\n4.Burrito\nElija el Numero de Opcion que deseen\n")
      	comida, _ := reader.ReadString('\n') // capturar string
        comida = strings.Replace(comida, "\n", "", -1) // capturar string
        comidaa, _ := strconv.Atoi(comida) /// convertimos el string a entero
        
        if comidaa==1 { // condicional pedido PIZZA
        	fmt.Print("\nQue Cantidad Deseas Agregar\n")
        	fmt.Scanf("%d\n", &surtidooP)
        	if surtidooP ==0 {
        	resultadoP, err := surtido(pizza, surtidooP)
   			 fmt.Printf("\nTotal de Pizzas en existencia = %d\n", resultadoP)
    		if err != nil {
        	fmt.Println(err)
    		}}
    		if surtidooP!=0 {
    			pizza=pizza+surtidooP
    			fmt.Print("\nComida surtida Con Exito = ", pizza,"\n")
    			
    		}
        }

        if comidaa==2 { // condicional pedido JOchos
        	fmt.Print("\nQue Cantidad Deseas Agregar\n")
        	fmt.Scanf("%d\n", &surtidooJ)
        	if surtidooJ ==0 {
        	resultadoJ, err := surtido(jochos, surtidooJ)
   			 fmt.Printf("\nTotal de JOchos en existencia = %d\n", resultadoJ)
    		if err != nil {
        	fmt.Println(err)
    		}}
    		if surtidooJ!=0 {
    			jochos=jochos+surtidooJ
    			fmt.Print("\nComida surtida Con Exito = ", jochos,"\n")
    			
    		}
        }

        if comidaa==3 { // condicional pedido Hamburgueza
        	fmt.Print("\nQue Cantidad Deseas Agregar\n")
        	fmt.Scanf("%d\n", &surtidooH)
        	if surtidooH ==0 {
        	resultadoH, err := surtido(hambuergueza, surtidooH)
   			 fmt.Printf("\nTotal de Hambuerguezas en existencia = %d\n", resultadoH)
    		if err != nil {
        	fmt.Println(err)
    		}}
    		if surtidooH!=0 {
    			hambuergueza=hambuergueza+surtidooH
    			fmt.Print("\nComida surtida Con Exito = ", hambuergueza,"\n")
    			
    		}
        }
        
        if comidaa==4 { // condicional pedido Burrito
        	fmt.Print("\nQue Cantidad Deseas Agregar\n")
        	fmt.Scanf("%d\n", &surtidooB)
        	if surtidooB ==0 {
        	resultadoB, err := surtido(burrito, surtidooB)
   			 fmt.Printf("\nTotal de Burritos en existencia = %d\n", resultadoB)
    		if err != nil {
        	fmt.Println(err)
    		}}
    		if surtidooB!=0 {
    			burrito=burrito+surtidooB
    			fmt.Print("\nComida surtida Con Exito = ", hambuergueza,"\n")
    			
    		}
        }

        if comidaa==4 { // condicional pedido Burrito
        	fmt.Print("\nQue Cantidad Deseas Agregar\n")
        	fmt.Scanf("%d\n", &surtidooB)
        	if surtidooB ==0 {
        	resultadoB, err := surtido(burrito, surtidooB)
   			 fmt.Printf("\nTotal de Burritos en existencia = %d\n", resultadoB)
    		if err != nil {
        	fmt.Println(err)
    		}}
    		if surtidooB!=0 {
    			burrito=burrito+surtidooB
    			fmt.Print("\nComida surtida Con Exito = ", hambuergueza,"\n")
    			
    		}
        }


    }//CIerre condicional 2
     /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    	if z_entero == 3 {// condicional DISPONIBLE

    		fmt.Print("\nRESTAURANTE GABYS Existencias\n")
    		fmt.Print("Total de Pizzas = ", pizza,"\n")
			fmt.Print("Total de Jochos = ", jochos,"\n")
			fmt.Print("Total de hambuerguezas = ", hambuergueza,"\n")
			fmt.Print("Total de burrito = ", burrito,"\n")

    	}// ciere OPCION 3

     /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    	if z_entero == 4 {// condicional salir
    		
    		break
    	}// ciere SALIR
    	 /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

       } //cierre for


         

}//cierro main