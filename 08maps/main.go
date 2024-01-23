package main

import "fmt"

func main() {
	fmt.Println("maps intro ")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println(languages)
	fmt.Println(languages["rb"])

	for key, value := range languages {
		fmt.Printf("For key  %v , value is %v \n" , key, value)

	}

}
