package main

import "fmt"

func main() {

	ratnesh := User{"Ratnesh", 25, true, "ratnesh@gmail.com"}
	fmt.Println(ratnesh)

	ratnesh.GetStatus()
	ratnesh.Newemail()
}

type User struct {
	Name   string
	Age    int
	Status bool
	Mail   string
}


func (u User ) GetStatus()  {
	fmt.Println(u.Name)
}

func (u User ) Newemail()  {
	u.Mail = "new@gmail.com"
	fmt.Println(u.Mail)
}