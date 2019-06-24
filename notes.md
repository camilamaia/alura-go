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
