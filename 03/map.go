package main

import "fmt"

func main() {
	animal := make(map[string]string)
	user := map[int]string{
		01: "ghalem",
		02: "filaly",
	}

	animal["lion"] = "leo"


	fmt.Println(user)
	fmt.Println(user[02])
	fmt.Println(animal)
	fmt.Println(animal["lion"])

}