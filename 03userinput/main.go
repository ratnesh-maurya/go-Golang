package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the course"
	println(welcome)

	raedline := bufio.NewReader(os.Stdin)
	// comma ok  , err err
	input, _ := raedline.ReadString('\n')
	fmt.Println(input)

}
