package main
import {
	"bufio"
	"fmt"
	"os"
	"strings"
}

func main(){
	reader := bufio.newReader(os.Stdin)
	fmt.Println("Primer Parcial")
	
	for{
		fmt.Print("Ingresa una palabra -> ")
		text, _ := reader.ReadString('\n')
		fmt.Print("Ingresa una vocal -> ")
		vocal, _ := reader.ReadString('\n')
		text = strings.replace(text, "a"||"e"||"i"||"o"||"u", vocal, -1)
		fmt.Println(text)
	}

	userFile := strings.toLower(text + ".txt")
		fout, err := os.Create(userFile)		
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		fout.WriteString(text)
}