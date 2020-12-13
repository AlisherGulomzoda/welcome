package main

import (
	"flag"
	"fmt"
	"strings"
)


func main() {
	flag.Parse()
	name := flag.Arg(0)
	code := flag.Arg(1)

	message := `Добро пожвлвать, {username}! Ваш код доступа: {code}.`
	message = strings.ReplaceAll(message, `{username}`, name)
	message = strings.ReplaceAll(message, `{code}`, code)
	
	fmt.Println(message)
}