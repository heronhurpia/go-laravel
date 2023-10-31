package main

import (
	"fmt"

	"github.com/heronhurpia/celeritas"
)

func main() {
	result := celeritas.TestFunc(3, 5)
	fmt.Println(result)

	result = celeritas.TestFunc2(8, 17)
	fmt.Println(result)

	result = celeritas.TestFunc3(4, 6)
	fmt.Println(result)

}
