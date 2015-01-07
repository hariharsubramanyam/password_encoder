package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var message = os.Args[1]
	var key, _ = strconv.ParseInt(os.Args[2], 10, 0)
	fmt.Println(message)
	fmt.Println(key)
}
