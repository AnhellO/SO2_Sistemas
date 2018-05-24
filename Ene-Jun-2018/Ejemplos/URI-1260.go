package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"sort"
	"strings"
)

func imprimeMapa(mapa map[string]int, nArboles int) {
	llaves := make([]string, 0, len(mapa))

	for llave := range mapa {
		llaves = append(llaves, llave)
	}

	sort.Strings(llaves)

	for _, llave := range llaves {
		fmt.Printf("%s %.4f\n", llave, (float64(mapa[llave]) * 100.0) / float64(nArboles))
	}

	return
}

func main() {

	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &N)
	arbol, err := reader.ReadString('\n')
	_ = arbol
	_ = err

	for caso := 1; caso <= N; caso++{
		mapa := make(map[string]int)
		nArboles := 0

		for {

			arbol, err := reader.ReadString('\n')

			if arbol == "\n" {
				imprimeMapa(mapa, nArboles)
				fmt.Println()
				break
			} else {
				arbol = strings.Replace(arbol, "\n", "", -1)
				mapa[arbol]++
				nArboles++
			}

			if err == io.EOF {
				imprimeMapa(mapa, nArboles)
				break
			}
		}

	}

}
