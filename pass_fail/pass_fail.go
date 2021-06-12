package main

import (
	"awesomeProject/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatalln(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "falling"
	}
	fmt.Println(status)

	keyboard.PrintConstant()

	fmt.Println(keyboard.Witcher)

}
