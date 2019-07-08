## Variables

```go
package main

import "fmt"
import "reflect"

func main() {
	name := "Camila" // var name string = "Camila"
	age := 27        // var age int = 27
	version := 1.1   // var version float64 = 1.1

	fmt.Println("Hello,", name)
	fmt.Println("Your age is", age)
	fmt.Println("Version of this program:", version)

	fmt.Println("Type of version:", reflect.TypeOf(version))
}
```

## Scan and Pointers
```go
package main

import "fmt"

func main() {
	name := "Camila"
	version := 1.1
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
```

## Flow control
If and Else:
```go
package main

import "fmt"

func main() {
	name := "Camila"
	version := 1.1
	fmt.Println("Hello,", name)
	fmt.Println("Version of this program:", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("The selected command was", command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing logs...")
	} else if command == 0 {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Invalid command")
	}
}
```

Switch:
```go
package main

import "fmt"

func main() {
	name := "Camila"
	version := 1.1
	fmt.Println("Hello,", name)
	fmt.Println("Version of this program:", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("The selected command was", command)

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid command")
	}
}
```

## Functions

```go
package main

import "fmt"
import "os"

func main() {
	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
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

	return command
}
```

## Functions with multiple returns

```go
package main

import "fmt"

func main() {
	name, age := nameAndAge()
	fmt.Println(name, "-", age, "years old")
}

func nameAndAge() (string, int) {
	name := "Camila"
	age := 27
	return name, age
}
```

How to ignore a return value:
```go
package main

import "fmt"

func main() {
	name, _ := nameAndAge()
	fmt.Println("My name is", name)
}

func nameAndAge() (string, int) {
	name := "Camila"
	age := 27
	return name, age
}
```

## For is Go's "while" and nett/http

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully")
	} else {
		fmt.Println("Site:", site, "is showing some problems. Status code:", resp.StatusCode)
	}
}
```
