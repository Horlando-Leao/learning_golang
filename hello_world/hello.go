package main

import (
	"fmt"
	"strings"
)

func main() {
	//declare variable
	var s strings.Builder

	//append string Hello
	s.WriteString("Oi! ")

	//append string from
	s.WriteString("Horlando, ")

	//append string educative
	s.WriteString("Como vai?")

	//Convert string builder to string and print it
	fmt.Println(s.String())
}
