package main

import "fmt"

var userDB = map[string]string{
	"saif": "lll",
}

func main() {
	if userDB["sif"] == "" {
		fmt.Println("not present")
		return
	}

	fmt.Println("preesent")
}
