package main

import (
	"flag"
	"fmt"
)

func main() {
	config := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	fmt.Println(*config)
}
