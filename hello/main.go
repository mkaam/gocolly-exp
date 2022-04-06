package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.Reverse("Amrullah"))
	fmt.Println(stringutil.ToUpper("Amrullah"))
}
