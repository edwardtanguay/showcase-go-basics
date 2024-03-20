package main

import (
	"fmt"

	"tanguay.eu/go/io/data"
)

var url = "http://google.com"

func main() {
	fmt.Println("hello from module")
	const pi = 3.14
	var message string
	var message2 = "hello"
	name := "test"
	price := 34
	isOpen := true

	println(pi, message, message2, price, isOpen)

	printData()
	fmt.Println(name001)
	fmt.Println(url)

	fmt.Println()

	fmt.Println(data.MaxSpeed)

	fmt.Println("Scores", data.Scores)

	for i := 0; i < len(data.Scores); i++ {}

	fmt.Printf("The name is: %s\n", name)


	// maps

	var codes map[int]bool

	fmt.Println("codes", codes)

	animals := map[int]string{90: "Dog", 91: "Cat", 92: "Cow", 93: "Bird", 94: "Rabbit"}
	fmt.Println("Map-2: ", animals)

	for key, value := range animals {
		fmt.Println(key, " -> ", value)
	}

	fmt.Printf("There are %d animals.\n", len(animals))

	// slices (any length, flexible)

	// not possible since animals is not a slice
	// animals := append(animals, {95: "Kangaroo"})

	data.Scores = append(data.Scores, 6)
	fmt.Println("data.Scores", data.Scores)

	names := []string {"one", "two"}
	names = append(names, "three")
	fmt.Printf("There are %d names.\n", len(names))

	nums := []int {1,2,3,4}
	nums = append(nums, 3)
	fmt.Printf("There are %d nums.\n", len(nums))
}