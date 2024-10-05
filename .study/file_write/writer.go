package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"/api/.study/file_write/test.txt",
		os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("waiwa\n")
	fmt.Println("Write file success!")
}
