package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			printLogs()
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
	sites := readSitesFromFile()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Something went wrong:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "is showing some problems. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	// returns an array of bytes. Good to be used when want to print the whole file
	// file, err := ioutil.ReadFile("sites.txt")

	file, err := os.Open("sites.txt") // returns the pointer

	if err != nil {
		fmt.Println("Something went wrong:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n') // single quote represents bytes
		if line == "" {
			break
		}
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	logFile.Close()
}

func printLogs() {
	logFile, err := ioutil.ReadFile("log.txt") // don't need to close the file since this is a higher level function

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(logFile))
}
