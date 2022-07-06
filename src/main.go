package main

import (
	"fmt"
	"github.com/guimassoqueto/gomod"
)

func main() {
	const h := helperType{
		name: "helper"
		completed: false
	}
	fmt.Println(h.name, h.completed)
}

