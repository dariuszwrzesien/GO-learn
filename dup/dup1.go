// Dup1 wyświetla tekst każdej linii, która pojawia się na standardowym wejściu więcej niż raz i poprzedza go liczbą wystąpień.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Println(counts)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s:\t%d\n", line, n)
			}
		}
	}

}
