package main

import (
	"fmt"

	helloGaes "github.com/hanifmaliki/go-module-test/pkg/hello-gaes"
	helloWorld "github.com/hanifmaliki/go-module-test/pkg/hello-world"
)

func main() {
	fmt.Println("Test")
	helloGaes.Print()
	helloWorld.Print()
}
