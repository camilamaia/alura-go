package main

import "fmt"

func main() {
	name := "Camila" // var name string = "Camila"
	version := 1.1   // var version float64 = 1.1
	fmt.Println("Hello,", name)
	fmt.Println("Version of this program:", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command) // fmt.Scanf("%d", &command)

	fmt.Println("Command pointer:", &command)
	fmt.Println("The selected command was", command)
}
