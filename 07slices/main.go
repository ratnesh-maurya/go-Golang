package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices intro ")
    var fruitlist  = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Fruit List is ", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println("Fruit List is ", fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println("Fruit List is ", fruitlist)

	highscores := make([] int , 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 458
	highscores[3] = 345 

	highscores = append(highscores, 555, 666, 777)

	fmt.Println("High Scores is ", highscores)
	sort.Ints(highscores)
	fmt.Println("High Scores is ", highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}