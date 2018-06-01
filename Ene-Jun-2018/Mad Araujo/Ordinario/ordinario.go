package main

import("bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "errors"
    "math")

func suma(valor1 int, valor2 int) (int, error) {
    if valor1==0 || valor2 == 0 {
        return 0, errors.New("No puedes sumar una letra")
    }

    return valor1 + valor2, nil;
}

func resta(valor1 int, valor2 int) (int, error) {
    if valor1 == 0 || valor2 == 0 {
        return 0, errors.New("No Tiene sentido Restar un 0")
    }

    return valor1 - valor2, nil;
}

func multiplicar(valor1 int, valor2 int) (int, error) {
    if valor1 == 0 || valor2 == 0 {
        return 0, errors.New("No Tiene sentido multiplicar por 0")
    }

    return valor1 * valor2, nil;
}

func division(valor1 int, valor2 int) (int, error) {
    if valor1 == 0 || valor2 == 0 {
        return 1, errors.New("No se puede dividir por 0")
    }

    return valor1 / valor2, nil;
}

func raiz(valor1 float64, valor2 float64) (float64, error) {
    if valor1 == 0 || valor2 <= 0 {
        return 0, errors.New("No se puede Negativo")
    }

    return math.Pow(valor1,1/valor2), nil;
}

func potencia(valor1 float64, valor2 float64) (float64, error) {
    if valor1 == 0 || valor2 == 0 {
        return 0, errors.New("No se puede potencia por 0")
    }

    return math.Pow(valor1,valor2), nil;
}





func main(){
	var as, ar, am, ad int //, 
	var bs, br, bm, bd int //, 
    var ara, ap float64
    var bra, bp float64

reader := bufio.NewReader(os.Stdin) // para capturar string
for{
fmt.Print("\nCALCULADORA FINAL \n1.SUMA\n2.RESTA\n3.MULTIPLICACION\n4.DIVISION\n5.RAIZ\n6.POTENCIA\n0.SALIR\n\n") // Opcion a elegir
        text, _ := reader.ReadString('\n') // capturar string
        text = strings.Replace(text, "\n", "", -1) // capturar string
       z_entero, _ := strconv.Atoi(text) /// convertimos el string a entero
        
        if z_entero == 1 {	///suma///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%d\n", &as)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%d\n", &bs)
        resultadoS, err := suma(as, bs)
        fmt.Printf("Resultado = %d\n", resultadoS)
        if err != nil {
        fmt.Println(err)
        }
    }
     
     if z_entero == 2 { ///resta///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%d\n", &ar)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%d\n", &br)
        resultadoR, err := resta(ar, br)
        fmt.Printf("Resultado = %d\n", resultadoR)
        if err != nil {
        fmt.Println(err)
        }
    }   

    if z_entero == 3 { ///Multiplicacion///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%d\n", &am)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%d\n", &bm)
        resultadoM, err := multiplicar(am, bm)
        fmt.Printf("Resultado = %d\n", resultadoM)
        if err != nil {
        fmt.Println(err)
        }
    }  


    if z_entero == 4 { ///Division///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%d\n", &ad)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%d\n", &bd)
        resultadoD, err := division(ad, bd)
        fmt.Printf("Resultado = %d\n", resultadoD)
        if err != nil {
        fmt.Println(err)
        }
    }   

    if z_entero == 5 { ///raiza///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%f\n", &ara)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%f\n", &bra)
        resultadoRa, err := raiz(ara, bra)
        f1:=float64(resultadoRa)
        fmt.Printf("Resultado = %f\n", f1)
        if err != nil {
        fmt.Println(err)
        }
    }   


    if z_entero == 6 { ///potencia///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        fmt.Print("\nINGRESE A\n")
            fmt.Scanf("%f\n", &ap)
        fmt.Print("\nINGRESE B\n")
            fmt.Scanf("%f\n", &bp)
        resultadoP, err := potencia(ap, bp)
        f:=float64(resultadoP)
        fmt.Printf("Resultado = %f\n", f)
        if err != nil {
        fmt.Println(err)
        }
    }   


    if z_entero == 0 { ///resta///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        break
    }   




     /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    	
    	 /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

       } //cierre for


         

}//cierro main