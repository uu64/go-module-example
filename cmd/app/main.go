package main

import (
	"fmt"

	"github.com/uu64/go-module-example/internal/greeting"
)

func main() {
	greeting.Hello()
	greeting.GoodBye()
	fmt.Println(greeting.AvailableLanguage())
}
