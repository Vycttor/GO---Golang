package main

import(
	"fmt"
	"bufio"
	"os"
	//"strconv"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Digite alguma coisa: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Você digitou: %q", input)

}