package main  
import("fmt")
func main() {
	var k,n,m,x,y int
	var resultado string
	fmt.Scanf("%d\n", &k)
	fmt.Scanf("%d %d\n", &n, &m)

       for i := 0; i <=k-1 ; i++ {
       	fmt.Scanf("%d %d\n", &x, &y)   
       	if(n==x || m==y){
       	resultado+="divisa"
       }    
       if(n<x && m<y){
       	resultado+="NE"
       }
       if(n>x && m<y){
       	resultado+="NO"
       }
       if(n>x && m>y){
       	resultado+="SO"
       }
       if(n<x && m>y){
       	resultado+="SE"
       }
       if (i==k-1) {
		fmt.Scanf("%d\n", &k)
        i=-1  
       if (k==0) {
       		break       		
       	}
       	fmt.Scanf("%d %d\n", &n, &m)
       }
   }
       	fmt.Println(resultado)      	
}