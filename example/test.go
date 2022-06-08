package main

import (
	"fmt"
	"strings"
)

func main() {
	// statusLangGetTemplate := make(map[int]map[string]string)
	// statusLangGetTemplate[1] =

	// fmt.Printf("33322:", statusLangGetTemplate)
	// f := fmt.Sprintf("html_%s", "3333")
	// fmt.Println("33322:", f)

	f := strings.Replace("oink \n oink \n oink", "\n", "<br/>", -1)
	fmt.Println(f)
	// fmt.Print("33322:", fmt.Sprintf("33322:", f))
}
