package main

import "fmt"
import "os"

func main() {
	var N, D int
	var alumnos [101]int
	
	fmt.Scanf("%d %d", &N, &D)
	
	for caso := 1; caso <= N; caso++ {
		for cena := 1; cena <= D ; cena++ {
			for xi := 1; xi <= N ; xi++ {
				var asistio int
				fmt.Scanf("%d", &asistio)
				if asistio == 1 {
					alumnos[xi]++;
				}
			}
		}
		
		var check bool
		for alumno := 1; alumno <= 100; alumno++ {
			if alumnos[alumno] == D {
				check = true
				fmt.Print("yes")
				break
			}
		}
		
		if check == false {
			fmt.Print("no")
		}
		
		fmt.Scanf("%d %d", &N, &D)
		if N == 0 && D == 0 {
			os.Exit(0)
		}
		
		fmt.Println()
	}
}