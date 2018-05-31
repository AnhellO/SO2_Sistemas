package main


	import (

		"fmt"
		"errors"
		
	)
	func division(val1 int, val2 int) (int, error) {
    if val1 >= 0 && val2 == 0 {
       return 0, errors.New("No te pases, M es 0!")
    }

    return val1 / val2, nil;
}


func main() {
        
    var N, M int
   
    for {

        fmt.Scanf("%d %d\n", &N, &M)
        resulDivision, err := division(N, M)     
                
    if N ==0 && M==0{
    	break

    }
    fmt.Printf("Resultado = %d\n", resulDivision)
    if err != nil {
    	fmt.Println(err)
    } 
   
  }

    

  
          
    
  
     
  

        
    
        
    
    
    
    
   


    
    








}
    

    