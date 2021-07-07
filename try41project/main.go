package main

import (
	greeting "../try41greeting" //一個上の階層に行ってtry41greetingpackage
	"fmt"
)

func main() {
	fmt.Println(greeting.Do())
}
