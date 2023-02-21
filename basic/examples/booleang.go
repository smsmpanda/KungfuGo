package main

import (
	"fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g.3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}

func main() {
	bool1 := true
	if bool1 {
		fmt.Println("The value is true")
	} else {
		fmt.Print("The value is false")
	}

	if val := 10; val > 5 {
		fmt.Println("value >= 5")
	} else {
		fmt.Println("value < 5")
	}
}