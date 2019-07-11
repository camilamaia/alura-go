package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorings = 5
const delay = 5

func main() {
	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func nameAndAge() (string, int) {
	name := "Camila"
	age := 27
	return name, age
}

func showIntroduction() {
	name := "Camila"
	version := 1.1
	fmt.Println("Hello,", name)
	fmt.Println("Version of this program:", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The selected command was", command)
	fmt.Println("")

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.loadsmart.com",
	}

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			checkSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func checkSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully")
	} else {
		fmt.Println("Site:", site, "is showing some problems. Status code:", resp.StatusCode)
	}
}
