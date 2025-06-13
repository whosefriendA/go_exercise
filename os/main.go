package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()

	fmt.Println(pid)

	gomod, err := os.Open("go.mod")

	defer gomod.Close()
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err == nil {
		fmt.Println(gomod)
	}

	info, err := os.Stat("go.mod")
	fmt.Println(info)

	os.Exit(1)
}
