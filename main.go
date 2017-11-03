package main

import (
	"fmt"
	"github.com/awaluddin/gomig/migration"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		first := os.Args[1]
		second := os.Args[2]
		Checking(first, second)
	} else {
		fmt.Println("gomig [first] [second]")
	}

}

func Checking(first, second string) {

	switch first {
	case "create":
		migration.Initialization()
	case "call":
		migration.CallAPI()
		break
	}

	fmt.Println(first + " " + second)
}
