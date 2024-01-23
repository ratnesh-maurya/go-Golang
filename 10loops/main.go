package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	// for i := range days {
	// fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("index is %v and value is %v \n", index, value)
	}

	roughvalue := 1
	for roughvalue < 10 {

		if roughvalue == 2{
			goto lco
		}
		if roughvalue == 5 {
			break 
		}
		fmt.Println(roughvalue)
		roughvalue++
	}

	lco:
	   fmt.Println("I am inside lco")
}
