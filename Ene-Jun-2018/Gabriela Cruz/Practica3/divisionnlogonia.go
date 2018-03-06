package main


	import (

		"fmt"
	)

	func main() {

        var k,x,y,n,m int
		var resultado string

		
		fmt.Scanf("%d\n", &k)
		fmt.Scanf("%d %d\n", &n, &m)
        
        for i:=0; i<=k-1; i++{
        	fmt.Scanf("%d %d\n", &x, &y)

        	if(x==n || y==m){
        		resultado+="DIVISA"
        	}
        	if(x>n && y>m){
        		resultado+="NE"
        	}
        	if(x<n && y>m){
        		resultado+="NO"
        	}
        	if(x<n && y<m){
        		resultado+="SO"
        	}
        	if(x>n && y<m){
        		resultado+="SE"
        	}
           	if(i==k-1){
           		fmt.Scanf("%d\n", &k)
        		
                i=-1;
                if(k==0){
                	break
                }
        		fmt.Scanf("%d %d\n", &n, &m)
        	} 
        	
        	
        }
        fmt.Println(resultado)
       
}