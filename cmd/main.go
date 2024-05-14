package main

import (
	"fmt"

	"github.com/AlucardMode/go-cli-todo/structures"
)

func main() {
	fmt.Println("Hello World")
	list := structures.List{Title: "Hello WOrld"}
	list.AddTask("New Task")
	list.AddTask("SecondTask")
	for _, v := range list.Tasks {
		fmt.Println(v.Description)
	}
}
