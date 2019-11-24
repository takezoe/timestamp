package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: timestamp <epoc_time>")
		os.Exit(1)
	}
	var epoc = os.Args[1]
	if len(epoc) == 7 {
		epoc = epoc + "000"
	}
	var i, _ = strconv.ParseInt(epoc, 10, 64)
	fmt.Println(i)
	var t = time.Unix(i, 0)
	fmt.Println(t)
}
